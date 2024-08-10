package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	body := getCotacao(ctx)
	createFile(body)

}

func getCotacao(ctx context.Context) []byte {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body

}

func createFile(body []byte) {
	f, err := os.Create("./client/cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(body)
}
