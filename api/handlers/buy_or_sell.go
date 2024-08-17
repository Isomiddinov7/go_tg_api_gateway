package handlers

import (
	"context"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/coins_service"

	"github.com/gin-gonic/gin"
)

// GetSell godoc
// @ID get_sell
// @Router /sell [POST]
// @Summary Get Sell
// @Description Get Sell
// @Tags Sell
// @Accept json
// @Produce json
// @Param profile body coins_service.BuyOrSellRequest true "BuyOrSellRequest"
// @Success 200 {object} http.Response{data=coins_service.BuyOrSellResponse} "SellBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSell(c *gin.Context) {

	var temp coins_service.BuyOrSellRequest

	err := c.ShouldBindJSON(&temp)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BuyOrSale().GetSell(
		context.Background(),
		&coins_service.BuyOrSellRequest{
			UserId:     temp.UserId,
			CoinId:     temp.CoinId,
			CoinAmount: temp.CoinAmount,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetBuy godoc
// @ID get_buy
// @Router /buy [POST]
// @Summary Get Buy
// @Description Get Buy
// @Tags Buy
// @Accept json
// @Produce json
// @Param profile body coins_service.BuyOrSellRequest true "BuyOrSellRequest"
// @Success 200 {object} http.Response{data=coins_service.BuyOrSellResponse} "SellBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBuy(c *gin.Context) {

	var temp coins_service.BuyOrSellRequest

	err := c.ShouldBindJSON(&temp)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	resp, err := h.services.BuyOrSale().GetBuy(
		context.Background(),
		&coins_service.BuyOrSellRequest{
			UserId:     temp.UserId,
			CoinId:     temp.CoinId,
			CoinAmount: temp.CoinAmount,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}


