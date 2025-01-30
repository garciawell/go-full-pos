package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/garciawell/go-challenge-auction/config/logger"
	"github.com/garciawell/go-challenge-auction/internal/entity/bid_entity"
	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BidRepository) FindBidByAuction(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auctionId": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error(fmt.Sprintf("Error finding bid by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error finding bid by auctionId %s", auctionId))
	}

	var bidEntityMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error decoding bid by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error decoding bid by auctionId %s", auctionId))
	}

	var bidEntities []bid_entity.Bid
	for _, bid := range bidEntityMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: time.Unix(bid.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (bd *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auctionId": auctionId}

	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})
	var bidEntityMongo BidEntityMongo
	err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo)
	if err != nil {
		logger.Error(fmt.Sprintf("Error finding winning bid by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error finding winning bid by auctionId %s", auctionId))
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil

}
