package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/garciawell/go-challenge-cloud-run/cmd/types"
	"github.com/garciawell/go-challenge-cloud-run/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/weather/cep/{cep}", handler)
	fmt.Println("Server running on port 8081")
	http.ListenAndServe(":8081", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	data := internal.GetAddressByCep(w, cep)

	if data.Localidade != "" {
		dataWeather, err := internal.GetWeatherByCity(data.Localidade)
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
