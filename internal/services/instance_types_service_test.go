package services_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/RHEnVision/provisioning-backend/internal/clients"
	clientStub "github.com/RHEnVision/provisioning-backend/internal/clients/stubs"
	"github.com/RHEnVision/provisioning-backend/internal/dao/stubs"
	"github.com/RHEnVision/provisioning-backend/internal/preload"
	"github.com/RHEnVision/provisioning-backend/internal/services"
	"github.com/RHEnVision/provisioning-backend/internal/testing/identity"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListInstanceTypesHandler(t *testing.T) {
	t.Run("with region", func(t *testing.T) {
		var names []string
		ctx := stubs.WithAccountDaoOne(context.Background())
		ctx = identity.WithTenant(t, ctx)
		ctx = clientStub.WithSourcesClient(ctx)
		ctx = clientStub.WithEC2Client(ctx)

		rctx := chi.NewRouteContext()
		ctx = context.WithValue(ctx, chi.RouteCtxKey, rctx)
		rctx.URLParams.Add("ID", "1")
		req, err := http.NewRequestWithContext(ctx, "GET", "/api/provisioning/sources/1/instance_types?region=us-east-1", nil)
		require.NoError(t, err, "failed to create request")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(services.ListInstanceTypes)
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code, "Handler returned wrong status code")

		var result []clients.InstanceType

		err = json.NewDecoder(rr.Body).Decode(&result)
		require.NoError(t, err, "failed to decode response body")

		assert.Equal(t, 3, len(result), "expected three result in response json")
		for _, it := range result {
			names = append(names, it.Name.String())
		}
		assert.Contains(t, names, "a1.2xlarge", "expected result to contain a1.2xlarge instance type")
		assert.Contains(t, names, "c5.xlarge", "expected result to contain c5.xlarge instance type")
	})

	t.Run("without region", func(t *testing.T) {
		ctx := stubs.WithAccountDaoOne(context.Background())
		ctx = identity.WithTenant(t, ctx)
		ctx = clientStub.WithSourcesClient(ctx)
		ctx = clientStub.WithEC2Client(ctx)

		req, err := http.NewRequestWithContext(ctx, "GET", "/api/provisioning/sources/1/instance_types", nil)
		require.NoError(t, err, "failed to create request")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(services.ListInstanceTypes)
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code, "Handler returned wrong status code")

		assert.Contains(t, rr.Body.String(), "parameter is missing")
	})
}

func TestListAzureBuiltinInstanceTypesHandler(t *testing.T) {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/provisioning/v1/instance_types/azure", nil)
	require.NoError(t, err, "failed to create request")
	req.URL.RawQuery = url.Values{"region": {"westus2"}, "zone": {"1"}}.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.ListBuiltinInstanceTypes(preload.AzureInstanceType.InstanceTypesForZone))
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code, "Handler returned wrong status code")

	var result []clients.InstanceType
	err = json.NewDecoder(rr.Body).Decode(&result)
	require.NoError(t, err, "failed to decode response body")

	assert.Less(t, 1, len(result), "the instance types response is empty")
}
