package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/garciawell/01-challenge-labs-observability/cmd/types"
	"github.com/garciawell/01-challenge-labs-observability/utils"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Webserver struct {
	TemplateData *TemplateData
}

type CEPInputDTO struct {
	Cep string `json:"cep"`
}

// NewServer creates a new server instance
func NewServer(templateData *TemplateData) *Webserver {
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

	router.Post("/weather/cep", we.HandleRequest)
	return router
}

type TemplateData struct {
	Title              string
	BackgroundColor    string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}

func (h *Webserver) HandleRequest(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	ctx, span := h.TemplateData.OTELTracer.Start(ctx, h.TemplateData.RequestNameOTEL)
	defer span.End()

	time.Sleep(time.Microsecond * h.TemplateData.ResponseTime)

	var cepInput CEPInputDTO
	err := json.NewDecoder(r.Body).Decode(&cepInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if utils.IsString(cepInput.Cep) == false {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if len(cepInput.Cep) != 8 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		client := &http.Client{}
		url := h.TemplateData.ExternalCallURL + "/weather/cep/" + cepInput.Cep

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		req.Header.Add("key", "53cf7ef523ac48e5a1c02131251401")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao realizar a request %v\n", err)
		}
		defer resp.Body.Close()
		res, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao Ler a resposta %v\n", err)
		}
		var data types.ResponseApi
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao realizar o parse %v\n", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(types.ResponseApi{
			TempC: data.TempC,
			TempF: data.TempF,
			TempK: data.TempC + 273,
		})

	}
}
