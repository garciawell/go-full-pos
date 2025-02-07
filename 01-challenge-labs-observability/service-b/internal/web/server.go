package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/garciawell/01-challenge-labs-observability/service-b/internal"
	"github.com/garciawell/01-challenge-labs-observability/types"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Webserver struct {
	TemplateData *types.TemplateData
}

type CEPInputDTO struct {
	Cep string `json:"cep"`
}

// NewServer creates a new server instance
func NewServer(templateData *types.TemplateData) *Webserver {
	return &Webserver{
		TemplateData: templateData,
	}
}

// createServer creates a new server instance with go chi router
func (we *Webserver) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/weather/cep/{cep}", we.HandleRequest)
	return router
}

func (h *Webserver) HandleRequest(w http.ResponseWriter, r *http.Request) {

	// OTL
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	ctx, span := h.TemplateData.OTELTracer.Start(ctx, h.TemplateData.RequestNameOTEL)
	defer span.End()

	time.Sleep(time.Microsecond * h.TemplateData.ResponseTime)

	cep := chi.URLParam(r, "cep")
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))
	data := internal.GetAddressByCep(ctx, w, cep)

	if data.Localidade != "" {
		dataWeather, err := internal.GetWeatherByCity(ctx, data.Localidade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(types.ResponseApi{
			TempC: dataWeather.Current.TempC,
			TempF: dataWeather.Current.TempF,
			TempK: dataWeather.Current.TempC + 273,
		})
	}
}
