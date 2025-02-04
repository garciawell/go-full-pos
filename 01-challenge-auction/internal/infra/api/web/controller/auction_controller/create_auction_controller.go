package auction_controller

import (
	"context"
	"net/http"

	"github.com/garciawell/go-challenge-auction/config/rest_err"
	"github.com/garciawell/go-challenge-auction/internal/infra/api/web/validation"
	"github.com/garciawell/go-challenge-auction/internal/usecase/auction_usecase"
	"github.com/gin-gonic/gin"
)

type auctionController struct {
	auctionUseCase auction_usecase.AuctionUseCase
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCase) *auctionController {
	return &auctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (a *auctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := a.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
