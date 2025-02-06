package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/garciawell/go-challenge-cloud-run/cmd/types"
	"github.com/garciawell/go-challenge-cloud-run/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CEPInputDTO struct {
	Cep string `json:"cep"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/weather/cep", handler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
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
		url := "http://localhost:8081/weather/cep/" + cepInput.Cep

		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("key", "53cf7ef523ac48e5a1c02131251401")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
