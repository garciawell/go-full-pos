package bid_usecase

import (
	"context"

	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
)

func (bu *BidUseCase) FindBidByAuction(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {
	bidList, err := bu.BidRepository.FindBidByAuction(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	var bidOutputDTOList []BidOutputDTO

	for _, bid := range bidList {
		bidOutputDTOList = append(bidOutputDTOList, BidOutputDTO{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}
	return bidOutputDTOList, nil
}

func (bu *BidUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError) {
	bid, err := bu.BidRepository.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	return &BidOutputDTO{
		Id:        bid.Id,
		UserId:    bid.UserId,
		AuctionId: bid.AuctionId,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}, nil
}
