package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Cep interface {
	SetName(name string)
}

type CepVia struct {
	Name        string `json:"name"`
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

type CepBrasil struct {
	Name         string `json:"name"`
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (c *CepVia) SetName(name string) {
	c.Name = name
}

func (c *CepBrasil) SetName(name string) {
	c.Name = name
}

func formatReturn(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getCep[T Cep](url string, name string, cn chan<- T) {
	resp, err := http.Get(url)
	body, err := formatReturn(resp, err)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	var c T
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Println("Error unmarshalling data:", err)
		return
	}

	c.SetName(name)
	cn <- c
	close(cn)
}

func main() {
	cn1 := make(chan *CepVia)
	cn2 := make(chan *CepBrasil)
	cep := "01153000"
	go getCep[*CepVia]("https://viacep.com.br/ws/"+cep+"/json/", "ViaCep", cn1)
	go getCep[*CepBrasil]("https://brasilapi.com.br/api/cep/v1/"+cep+"", "BrasilApi", cn2)

	select {
	case res := <-cn1:
		fmt.Printf("Nome %s\n", res.Name)
		fmt.Printf("Rua: %s, Bairro: %s, Cidade: %s, Estado: %s\n", res.Logradouro, res.Bairro, res.Localidade, res.Uf)
	case res := <-cn2:
		fmt.Printf("Nome %s\n", res.Name)
		fmt.Printf("Rua: %s, Bairro: %s, Cidade: %s, Estado: %s\n", res.Street, res.Neighborhood, res.City, res.State)
	case <-time.After(time.Second):
		println("Timeout")
	}
}
