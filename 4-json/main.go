package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance" validate:"gt=0"`
}

func main() {
	account := Account{Number: 1, Balance: 100}

	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(account)

	pureJson := []byte(`{"n":2,"s":200}`)
	var accountX Account

	err = json.Unmarshal(pureJson, &accountX)
	if err != nil {
		panic(err)
	}
	println(accountX.Balance)
}
