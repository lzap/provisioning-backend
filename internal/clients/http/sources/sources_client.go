package sources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	stdhttp "net/http"
	"net/url"
	"strings"

	"github.com/RHEnVision/provisioning-backend/internal/cache"
	"github.com/RHEnVision/provisioning-backend/internal/clients"
	"github.com/RHEnVision/provisioning-backend/internal/clients/http"
	"github.com/RHEnVision/provisioning-backend/internal/config"
	"github.com/RHEnVision/provisioning-backend/internal/headers"
	"github.com/RHEnVision/provisioning-backend/internal/models"
	"github.com/RHEnVision/provisioning-backend/internal/ptr"
	"github.com/RHEnVision/provisioning-backend/internal/telemetry"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
)

const TraceName = telemetry.TracePrefix + "internal/clients/http/sources"

type sourcesClient struct {
	client *ClientWithResponses
}

func init() {
	clients.GetSourcesClient = newSourcesClient
}

func logger(ctx context.Context) zerolog.Logger {
	return zerolog.Ctx(ctx).With().Str("client", "sources").Logger()
}

func newSourcesClient(ctx context.Context) (clients.Sources, error) {
	return NewSourcesClientWithUrl(ctx, config.Sources.URL)
}

// NewSourcesClientWithUrl allows customization of the URL for the underlying client.
// It is meant for testing only, for production please use clients.GetSourcesClient.
func NewSourcesClientWithUrl(ctx context.Context, url string) (clients.Sources, error) {
	c, err := NewClientWithResponses(url, func(c *Client) error {
		c.Client = http.NewPlatformClient(ctx, config.Sources.Proxy.URL)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &sourcesClient{client: c}, nil
}

type appType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type dataElement struct {
	Data []appType `json:"data"`
}

func (c *sourcesClient) Ready(ctx context.Context) error {
	ctx, span := otel.Tracer(TraceName).Start(ctx, "Ready")
	defer span.End()

	logger := logger(ctx)
	resp, err := c.client.ListApplicationTypes(ctx, &ListApplicationTypesParams{}, headers.AddSourcesIdentityHeader, headers.AddEdgeRequestIdHeader)
	if err != nil {
		logger.Error().Err(err).Msg("Readiness request failed for sources")
		return err
	}
	defer resp.Body.Close()

	err = http.HandleHTTPResponses(ctx, resp.StatusCode)
	if err != nil {
		return fmt.Errorf("ready call: %w", err)
	}
	return nil
}

func (c *sourcesClient) ListProvisioningSourcesByProvider(ctx context.Context, provider models.ProviderType) ([]*clients.Source, error) {
	logger := logger(ctx)
	ctx, span := otel.Tracer(TraceName).Start(ctx, "ListProvisioningSourcesByProvider")
	defer span.End()

	appTypeId, err := c.GetProvisioningTypeId(ctx)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get provisioning type id")
		return nil, fmt.Errorf("failed to get provisioning app type: %w", err)
	}

	sourcesProviderName := provider.SourcesProviderName()
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get provider name according to sources service")
		return nil, fmt.Errorf("failed to get provider name according to sources service: %w", err)
	}

	resp, err := c.client.ListApplicationTypeSourcesWithResponse(ctx, appTypeId, &ListApplicationTypeSourcesParams{}, headers.AddSourcesIdentityHeader,
		headers.AddEdgeRequestIdHeader, BuildQuery("filter[source_type][name]", sourcesProviderName))
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to fetch ApplicationTypes from sources")
		return nil, fmt.Errorf("failed to get ApplicationTypes: %w", err)
	}

	err = http.HandleHTTPResponses(ctx, resp.StatusCode())
	if err != nil {
		if errors.Is(err, clients.NotFoundErr) {
			return nil, fmt.Errorf("list provisioning sources call: %w", http.SourceNotFoundErr)
		}
		return nil, fmt.Errorf("list provisioning sources call: %w", err)
	}

	result := make([]*clients.Source, 0, len(*resp.JSON200.Data))

	for _, src := range *resp.JSON200.Data {
		newSrc := clients.Source{
			ID:           ptr.From(src.Id),
			Name:         ptr.From(src.Name),
			SourceTypeID: ptr.From(src.SourceTypeId),
			Uid:          ptr.From(src.Uid),
		}
		result = append(result, &newSrc)
	}

	return result, nil
}

func (c *sourcesClient) ListAllProvisioningSources(ctx context.Context) ([]*clients.Source, error) {
	logger := logger(ctx)
	ctx, span := otel.Tracer(TraceName).Start(ctx, "ListAllProvisioningSources")
	defer span.End()

	appTypeId, err := c.GetProvisioningTypeId(ctx)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get provisioning type id")
		return nil, fmt.Errorf("failed to get provisioning app type: %w", err)
	}

	resp, err := c.client.ListApplicationTypeSourcesWithResponse(ctx, appTypeId, &ListApplicationTypeSourcesParams{}, headers.AddSourcesIdentityHeader, headers.AddEdgeRequestIdHeader)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to fetch ApplicationTypes from sources")
		return nil, fmt.Errorf("failed to get ApplicationTypes: %w", err)
	}

	err = http.HandleHTTPResponses(ctx, resp.StatusCode())
	if err != nil {
		if errors.Is(err, clients.NotFoundErr) {
			return nil, fmt.Errorf("list provisioning sources call: %w", http.SourceNotFoundErr)
		}
		return nil, fmt.Errorf("list provisioning sources call: %w", err)
	}

	result := make([]*clients.Source, len(*resp.JSON200.Data))
	for i, src := range *resp.JSON200.Data {
		newSrc := clients.Source{
			ID:           ptr.From(src.Id),
			Name:         ptr.From(src.Name),
			SourceTypeID: ptr.From(src.SourceTypeId),
			Uid:          ptr.From(src.Uid),
		}
		result[i] = &newSrc
	}
	return result, nil
}

func (c *sourcesClient) GetAuthentication(ctx context.Context, sourceId string) (*clients.Authentication, error) {
	logger := logger(ctx)
	ctx, span := otel.Tracer(TraceName).Start(ctx, "GetAuthentication")
	defer span.End()

	// Get all the authentications linked to a specific source
	resp, err := c.client.ListSourceAuthenticationsWithResponse(ctx, sourceId, &ListSourceAuthenticationsParams{}, headers.AddSourcesIdentityHeader,
		headers.AddEdgeRequestIdHeader, BuildQuery("filter[resource_type]", "Application", "filter[authtype][starts_with]", "provisioning"))
	if err != nil {
		return nil, fmt.Errorf("cannot list provisioning source authentication of type application: %w", err)
	}

	err = http.HandleHTTPResponses(ctx, resp.StatusCode())
	if err != nil {
		if errors.Is(err, clients.NotFoundErr) {
			return nil, fmt.Errorf("get source authentication call: %w", http.AuthenticationForSourcesNotFoundErr)
		}
		return nil, fmt.Errorf("get source authentication call: %w", err)
	}

	if len(*resp.JSON200.Data) != 0 {
		auth := (*resp.JSON200.Data)[0]
		authentication := clients.NewAuthenticationFromSourceAuthType(ctx, *auth.Username, string(*auth.Authtype), *auth.ResourceId)
		return authentication, nil
	} else {
		logger.Trace().Msgf("Source does not have provisioning authentications of type application")
		return nil, clients.MissingProvisioningSources
	}
}

func (c *sourcesClient) GetProvisioningTypeId(ctx context.Context) (string, error) {
	appTypeId, err := cache.FindAppTypeId(ctx)
	if errors.Is(err, cache.ErrNotFound) {
		appTypeId, err = c.loadAppId(ctx)
		if err != nil {
			return "", err
		}
		err = cache.SetAppTypeId(ctx, appTypeId)
		if err != nil {
			return "", fmt.Errorf("unable to store app type id to cache: %w", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("unable to get app type id from cache: %w", err)
	}

	return appTypeId, nil
}

func (c *sourcesClient) loadAppId(ctx context.Context) (string, error) {
	logger := logger(ctx)
	logger.Trace().Msg("Fetching the Application Type ID of Provisioning for Sources")

	resp, err := c.client.ListApplicationTypes(ctx, &ListApplicationTypesParams{}, headers.AddSourcesIdentityHeader, headers.AddEdgeRequestIdHeader)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to fetch ApplicationTypes from sources")
		return "", fmt.Errorf("failed to fetch ApplicationTypes: %w", err)
	}
	defer resp.Body.Close()

	err = http.HandleHTTPResponses(ctx, resp.StatusCode)
	if err != nil {
		if errors.Is(err, clients.NotFoundErr) {
			return "", fmt.Errorf("load app ID call: %w", http.ApplicationTypeNotFoundErr)
		}
		return "", fmt.Errorf("load app ID call: %w", err)
	}

	var appTypesData dataElement
	if err = json.NewDecoder(resp.Body).Decode(&appTypesData); err != nil {
		return "", fmt.Errorf("could not unmarshal application type response: %w", err)
	}
	for _, t := range appTypesData.Data {
		if t.Name == "/insights/platform/provisioning" {
			logger.Trace().Msgf("The Application Type ID found: '%s' and it got cached", t.Id)
			return t.Id, nil
		}
	}
	return "", http.ApplicationTypeNotFoundErr
}

func BuildQuery(keysAndValues ...string) func(ctx context.Context, req *stdhttp.Request) error {
	return func(ctx context.Context, req *stdhttp.Request) error {
		if len(keysAndValues)%2 != 0 {
			return http.NotEvenErr
		}
		queryParams := make([]string, 0)
		for i := 0; i < len(keysAndValues); i += 2 {
			key := url.QueryEscape(keysAndValues[i])
			value := url.QueryEscape(keysAndValues[i+1])
			queryParams = append(queryParams, fmt.Sprintf("%s=%s", key, value))
		}

		req.URL.RawQuery = strings.Join(queryParams, "&")
		return nil
	}
}
