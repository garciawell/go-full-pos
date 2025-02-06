package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/garciawell/go-challenge-cloud-run/cmd/types"
	"github.com/garciawell/go-challenge-cloud-run/utils"
)

func GetWeatherByCity(city string) (types.Weather, error) {
	client := &http.Client{}
	removeAccent := utils.RemoveAccents(city)
	encodedCity := url.QueryEscape(removeAccent)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?q=%s", encodedCity)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("key", "53cf7ef523ac48e5a1c02131251401")
	if err != nil {
		return types.Weather{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return types.Weather{}, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao Ler a resposta %v\n", err)
	}
	var data types.Weather
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao realizar o parse %v\n", err)
	}
	return data, nil
}
