package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição %v\n", err)
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

		fmt.Println(data.Logradouro)

		file, err := os.Create("./5-cep/rua.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao Criar arquivo %v\n", err)
		}

		defer file.Close()
		file.WriteString(fmt.Sprintf("Rua: %s, Estado: %s", data.Logradouro, data.Uf))
	}
}
