package mongodb

import (
	"context"
	"os"

	"github.com/garciawell/go-challenge-auction/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConn(ctx context.Context) (*mongo.Database, error) {
	// Connect to MongoDB
	mongoURL := os.Getenv(MONGODB_URL)
	mongoDatabase := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Error("Error connecting to MongoDB", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error pinging MongoDB", err)
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}
