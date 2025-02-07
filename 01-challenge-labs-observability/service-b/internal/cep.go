package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/garciawell/01-challenge-labs-observability/types"
)

func GetAddressByCep(ctx context.Context, w http.ResponseWriter, cep string) types.CEP {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")

	if len(cep) != 8 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid zipcode"})
		return types.CEP{}
	}

	if req.StatusCode == 400 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot find zipcode"})
		return types.CEP{}
	}

	if err != nil {
		return types.CEP{}
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao Ler a resposta %v\n", err)
	}

	var data types.CEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao realizar o parse %v\n", err)
	}

	if data.Localidade == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Endereço não encontrado"})
		return types.CEP{}
	}

	return data
}
