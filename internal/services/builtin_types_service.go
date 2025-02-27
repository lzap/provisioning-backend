package services

import (
	"net/http"
	"strings"
	"time"

	"github.com/RHEnVision/provisioning-backend/internal/clients"
	"github.com/RHEnVision/provisioning-backend/internal/payloads"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

type InstanceTypesForZoneFunc func(region, zone string, supported *bool) ([]*clients.InstanceType, error)

func ListBuiltinInstanceTypes(typeFunc InstanceTypesForZoneFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		region := strings.ToLower(r.URL.Query().Get("region"))
		zone := strings.ToLower(r.URL.Query().Get("zone"))
		supported, err := ParseBool(r.URL.Query().Get("supported"))
		if err != nil {
			renderError(w, r, payloads.NewInvalidRequestError(r.Context(), "parameter 'supported' could not be parsed", err))
			return
		}

		if region == "" {
			renderError(w, r, payloads.NewMissingRequestParameterError(r.Context(), "region parameter is missing"))
			return
		}

		start := time.Now()
		instances, err := typeFunc(region, zone, supported)
		logger := zerolog.Ctx(r.Context())
		logger.Trace().TimeDiff("duration", time.Now(), start).Msg("Listed instance types")
		if err != nil {
			renderError(w, r, payloads.NewInvalidRequestError(r.Context(), "instance types not found for selected region and zone", err))
			return
		}

		if err := render.RenderList(w, r, payloads.NewListInstanceTypeResponse(instances)); err != nil {
			renderError(w, r, payloads.NewRenderError(r.Context(), "unable to render instance types list", err))
			return
		}
	}
}
