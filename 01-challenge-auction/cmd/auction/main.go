package main

import (
	"context"
	"log"

	"github.com/garciawell/go-challenge-auction/config/database/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	_, err := mongodb.NewMongoDBConn(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
