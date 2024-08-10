package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	body, err := getCotacao(ctx)
	if err != nil {
		panic(err)
	}
	createFile(body)
}

func getCotacao(ctx context.Context) (body []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	return body, err
}

func createFile(body []byte) {
	f, err := os.Create("./cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Arquivo criado com sucesso...")
	f.Write([]byte("Bid: " + string(body)))
}
