package bid_entity

import (
	"context"
	"time"

	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

type BidEntityRepository interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
	FindBidByAuction(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}
