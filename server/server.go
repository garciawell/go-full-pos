package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var port = "8080"

type UsdBrl struct {
	Usdbrl struct {
		Code       string `json:"-"`
		Codein     string `json:"-"`
		Name       string `json:"-"`
		High       string `json:"-"`
		Low        string `json:"-"`
		VarBid     string `json:"-"`
		PctChange  string `json:"-"`
		Bid        string `json:"bid"`
		Ask        string `json:"-"`
		Timestamp  string `json:"-"`
		CreateDate string `json:"-"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", handleCurrency)

	if _, err := os.Stat("./cotacao.db"); err != nil {
		fmt.Printf("File not exists\n")
		fmt.Printf("Creating file")
		f, err := os.Create("./cotacao.db")
		if err != nil {
			panic(err)
		}
		fmt.Println(f)
	}

	db, err := sql.Open("sqlite3", "./cotacao.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("Banco conectado com sucesso...")
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	fmt.Println("Listening port " + port + " ...")
	http.ListenAndServe("localhost:"+port, nil)
}

func handleCurrency(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	data, err := getCurrency(ctx)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

}

func getCurrency(rx context.Context) (data *UsdBrl, err error) {
	req, err := http.NewRequestWithContext(rx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var currency UsdBrl

	format, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(format, &currency)
	if err != nil {
		panic(err)
	}

	return &currency, err
}
