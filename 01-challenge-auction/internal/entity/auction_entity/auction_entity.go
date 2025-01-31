package auction_entity

import (
	"context"
	"time"

	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
	"github.com/google/uuid"
)

func CreateAuction(productName, category, description string, condition ProductCondition) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		Id:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	err := auction.Validate()
	if err != nil {
		return nil, err
	}

	return auction, nil
}

func (au *Auction) Validate() *internal_error.InternalError {
	if len(au.ProductName) <= 1 ||
		len(au.Category) <= 2 ||
		(len(au.Description) <= 10 && (au.Condition != New && au.Condition != Used && au.Condition != Refurbished)) {
		return internal_error.NewBadRequestError("Invalid auction data")
	}

	return nil
}

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductCondition int
type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

const (
	New ProductCondition = iota
	Used
	Refurbished
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction *Auction) *internal_error.InternalError
	FindAuctionById(ctx context.Context, id string) (*Auction, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]Auction, *internal_error.InternalError)
}
