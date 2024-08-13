package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var port = "8080"

type UsdBrl struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

var db *sql.DB

func init() {
	if _, err := os.Stat("./cotacao.db"); err != nil {
		fmt.Println("Creating DB")
		f, err := os.Create("./cotacao.db")
		if err != nil {
			panic(err)
		}
		fmt.Println(f)
	}

	conn, err := sql.Open("sqlite3", "./cotacao.db")
	if err != nil {
		log.Println("TESTE")
		panic(err)
	}
	db = conn
	fmt.Println("Banco conectado com sucesso...")

	sql := `
	CREATE TABLE IF NOT EXISTS currency (id integer not null primary key, 
	code text, 
	codein text, 
	name text, 
	hight text, 
	low text,
	varBid text,
	pctChange text,
	bid text,
	ask text,
	timestamp text,
	create_date text
	);
	`
	_, err = db.Exec(sql)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sql)
		return
	}
}

func main() {
	http.HandleFunc("/cotacao", handleCurrency)
	fmt.Println("Listening port " + port + " ...")
	defer db.Close()
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

	insertDB(db, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Usdbrl.Bid)

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

func insertDB(db *sql.DB, currency *UsdBrl) {
	ctx := context.Background()
	ctx, close := context.WithTimeout(ctx, time.Millisecond*10)
	defer close()
	log.Println("Inserting record ...")
	insertDB := `INSERT INTO currency(code, codeIn, name, hight, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	data, err := db.PrepareContext(ctx, insertDB)
	if err != nil {
		fmt.Println(err)
	}

	_, err = data.Exec(currency.Usdbrl.Code,
		currency.Usdbrl.Codein,
		currency.Usdbrl.Name,
		currency.Usdbrl.High,
		currency.Usdbrl.Low,
		currency.Usdbrl.VarBid,
		currency.Usdbrl.PctChange,
		currency.Usdbrl.Bid,
		currency.Usdbrl.Ask,
		currency.Usdbrl.Timestamp,
		currency.Usdbrl.CreateDate)
	if err != nil {
		fmt.Println(err)
	}
}
