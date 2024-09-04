package handlers

import (
	"context"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/users_service"
	"go_tg_api_gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserSell godoc
// @ID User_sell
// @Router /user/sell [POST]
// @Summary UserSell
// @Description  UserSell
// @Tags User
// @Accept json
// @Produce json
// @Param file formData file true "Upload file"
// @Param userId formData string true "User ID"
// @Param coinId formData string true "Coin ID"
// @Param coin_amount formData string true "CoinAmount"
// @Param message formData string true "Message"
// @Param card_number formData string true "Card Number"
// @Success 200 {object} http.Response{data=string} "GetUserSellBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UserSell(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		h.handleResponse(c, http.BadRequest, gin.H{"error": "Unable to get file"})
		return
	}
	imageURL, err := utils.UploadImage(file)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	var user_sell users_service.UserSellRequest

	user_sell.UserId = c.PostForm("userId")
	user_sell.CoinId = c.PostForm("coinId")
	user_sell.CoinAmount = c.PostForm("coin_amount")
	user_sell.Message = c.PostForm("message")
	user_sell.CardNumber = c.PostForm("card_number")
	user_sell.CheckImg = imageURL

	_, err = h.services.UserTransaction().UserSell(
		c.Request.Context(),
		&user_sell,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, gin.H{"message": "Transaction created successfully"})
}

// UserBuy godoc
// @ID User_buy
// @Router /user/buy [POST]
// @Summary UserBuy
// @Description  UserBuy
// @Tags User
// @Accept json
// @Produce json
// @Param file formData file true "Upload file"
// @Param userId formData string true "User ID"
// @Param coinId formData string true "Coin ID"
// @Param address formData string true "Address"
// @Param coin_amount formData string true "CoinAmount"
// @Param message formData string true "Message"
// @Success 200 {object} http.Response{data=string} "GetUserBuyBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UserBuy(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		h.handleResponse(c, http.BadRequest, gin.H{"error": "Unable to get file"})
		return
	}
	imageURL, err := utils.UploadImage(file)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	var user_buy users_service.UserBuyRequest
	user_buy.UserId = c.PostForm("userId")
	user_buy.CoinId = c.PostForm("coinId")
	user_buy.CoinAmount = c.PostForm("coin_amount")
	user_buy.Message = c.PostForm("message")
	user_buy.Address = c.PostForm("address")
	user_buy.PayImg = imageURL

	_, err = h.services.UserTransaction().UserBuy(
		c.Request.Context(),
		&user_buy,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "ok")
}

// @Security ApiKeyAuth
// AllUserSell godoc
// @ID get_transaction_sell_list
// @Router /user/sell [GET]
// @Summary Get Transactions List
// @Description  Get Transactions List
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param status query string false "status"
// @Success 200 {object} http.Response{data=users_service.GetListUserSellTransactionResponse} "GetAllCoinResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AllUserSell(c *gin.Context) {

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.UserTransaction().AllUserSell(
		context.Background(),
		&users_service.GetListUserTransactionRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Status: c.Query("status"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// AllUserBuy godoc
// @ID get_transaction_by_list
// @Router /user/buy [GET]
// @Summary Get Transactions List
// @Description  Get Transactions List
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param status query string false "status"
// @Success 200 {object} http.Response{data=users_service.GetListUserBuyTransactionResponse} "GetAllCoinResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AllUserBuy(c *gin.Context) {

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.UserTransaction().AllUserBuy(
		context.Background(),
		&users_service.GetListUserTransactionRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Status: c.Query("status"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
