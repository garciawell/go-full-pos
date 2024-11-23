package main

import (
	"database/sql"
	"fmt"

	"github.com/garciawell/go-full-pos/39-DI/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a repository
	repository := product.NewProductRepository(db)

	// Create a useCase
	useCase := product.NewProductUseCase(repository)

	product, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
