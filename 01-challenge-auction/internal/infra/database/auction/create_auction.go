package auction

import (
	"context"

	"github.com/garciawell/go-challenge-auction/internal/entity/auction_entity"
	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(collection *mongo.Database) *AuctionRepository {
	return &AuctionRepository{Collection: collection.Collection("auctions")}
}

func (repo *AuctionRepository) CreateAuction(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {
	auctionMongo := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	_, err := repo.Collection.InsertOne(nil, auctionMongo)
	if err != nil {
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	return nil
}
