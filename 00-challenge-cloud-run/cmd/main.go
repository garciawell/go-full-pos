package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Weather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

type ResponseApi struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/weather/cep/{cep}", handler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	data := getAddressByCep(w, cep)

	if data.Localidade != "" {
		dataWeather, err := getWeatherByCity(data.Localidade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseApi{
			TempC: dataWeather.Current.TempC,
			TempF: dataWeather.Current.TempF,
			TempK: dataWeather.Current.TempC + 273,
		})
	}
}

func getAddressByCep(w http.ResponseWriter, cep string) CEP {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")

	if len(cep) != 8 {
		// JSON error response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": "CEP inválido"})
		return CEP{}
	}

	if req.StatusCode == 400 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "CEP não encontrado"})
		return CEP{}
	}

	if err != nil {
		return CEP{}
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao Ler a resposta %v\n", err)
	}

	var data CEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao realizar o parse %v\n", err)
	}

	if data.Localidade == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Endereço não encontrado"})
		return CEP{}
	}

	return data
}

func getWeatherByCity(city string) (Weather, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://api.weatherapi.com/v1/current.json?q=Curitiba", nil)
	req.Header.Add("key", "53cf7ef523ac48e5a1c02131251401")
	if err != nil {
		return Weather{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao Ler a resposta %v\n", err)
	}
	var data Weather
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao realizar o parse %v\n", err)
	}
	return data, nil
}
