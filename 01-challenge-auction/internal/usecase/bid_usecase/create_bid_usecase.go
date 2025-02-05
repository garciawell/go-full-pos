package bid_usecase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/garciawell/go-challenge-auction/config/logger"
	"github.com/garciawell/go-challenge-auction/internal/entity/bid_entity"
	internal_error "github.com/garciawell/go-challenge-auction/internal/internal_erro"
)

type BidInputDTO struct {
	Id        string  `json:"id"`
	UserId    string  `json:"user_id"`
	AuctionId string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02T15:04:05Z07:00"`
}

type BidUseCase struct {
	BidRepository bid_entity.BidEntityRepository

	timer         *time.Timer
	maxBatchSize  int
	batchInterval time.Duration
	bidChannel    chan bid_entity.Bid
}

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, inputDTO BidInputDTO) *internal_error.InternalError
	FindBidByAuction(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError)
}

func NewBidUseCase(bidRepository bid_entity.BidEntityRepository, maxBatchSize int, batchInterval time.Duration) *BidUseCase {
	bid := &BidUseCase{
		BidRepository: bidRepository,
		maxBatchSize:  getMaxBatchSize(),
		batchInterval: getMaxBatchSizeInterval(),
		timer:         time.NewTimer(getMaxBatchSizeInterval()),
		bidChannel:    make(chan bid_entity.Bid, getMaxBatchSize()),
	}

	bid.triggerCreateRoutine(context.Background())
	return bid
}

var bidBatch []bid_entity.Bid

func (b *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func() {
		defer close(b.bidChannel)

		for {
			select {
			case bidEntity, ok := <-b.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						if err := b.BidRepository.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("Error creating bid", err)
						}
					}
				}
				bidBatch = append(bidBatch, bidEntity)

				if len(bidBatch) >= b.maxBatchSize {
					if err := b.BidRepository.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("Error creating bid", err)
					}

					bidBatch = nil
					b.timer.Reset(b.batchInterval)
				}

			case <-b.timer.C: // Timer expired
				if err := b.BidRepository.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("Error creating bid", err)
				}
				bidBatch = nil
				b.timer.Reset(b.batchInterval)
			}
		}
	}()

}

func (b *BidUseCase) CreateBid(ctx context.Context, inputDTO BidInputDTO) *internal_error.InternalError {

	bidEntity, err := bid_entity.CreateBid(inputDTO.UserId, inputDTO.AuctionId, inputDTO.Amount)
	if err != nil {
		return err
	}

	b.bidChannel <- *bidEntity

	return nil
}

func getMaxBatchSizeInterval() time.Duration {
	batchInserInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInserInterval)
	if err != nil {
		return 3 * time.Minute
	}
	return duration
}

func getMaxBatchSize() int {
	maxBatchSize := os.Getenv("MAX_BATCH_SIZE")
	size, err := strconv.Atoi(maxBatchSize)
	if err != nil {
		return 5
	}
	return size
}
