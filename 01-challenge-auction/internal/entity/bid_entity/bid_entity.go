package bid_entity

import (
	"context"
	"time"

	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
	"github.com/google/uuid"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

func CreateBid(userId, auctionId string, amount float64) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		Id:        uuid.New().String(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	err := bid.Validate()
	if err != nil {
		return nil, err
	}
	return bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {
	err := uuid.Validate(b.UserId)
	if err != nil {
		return internal_error.NewBadRequestError("invalid user id")
	}
	err = uuid.Validate(b.AuctionId)
	if err != nil {
		return internal_error.NewBadRequestError("invalid auction id")
	}

	if b.Amount <= 0 {
		return internal_error.NewBadRequestError("Amount is not a valid value")
	}
	return nil
}

type BidEntityRepository interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
	FindBidByAuction(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}
