package main

import (
	"fmt"
	"net/http"

	"github.com/garciawell/go-full-pos/apis/configs"
	"github.com/garciawell/go-full-pos/apis/internal/entity"
	"github.com/garciawell/go-full-pos/apis/internal/infra/database"
	"github.com/garciawell/go-full-pos/apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var productHandler *handlers.ProductHandler

func init() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	productHandler = handlers.NewProductHandler(productDB)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)

	fmt.Println("Server is running on port 8000...")
	http.ListenAndServe(":8000", r)

}
