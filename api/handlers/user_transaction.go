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
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Param userId formData string true "User ID"
// @Param coinId formData string true "Coin ID"
// @Param coin_amount formData string true "CoinAmount"
// @Param message formData string true "Message"
// @Param card_number formData string true "Card Number"
// @Param card_holder_name formData string true "Card Holder Name"
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
	user_sell.CardHolderName = c.PostForm("card_holder_name")
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
// @Accept multipart/form-data
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
// @Param search query string false "search"
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
			Search: c.Query("search"),
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
// @Param search query string false "search"
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
			Search: c.Query("search"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// TransactionUpdate godoc
// @ID transaction_update
// @Router /transaction/{id} [PUT]
// @Summary TransactionUpdate
// @Description TransactionUpdate
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateTransaction true "UpdateTransactionBody"
// @Success 200 {object} http.Response{data=string} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) TransactionUpdate(c *gin.Context) {

	var data users_service.UpdateTransaction

	data.Id = c.Param("id")

	if !utils.IsValidUUID(data.Id) {
		h.handleResponse(c, http.InvalidArgument, "User id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.services.UserTransaction().TransactionUpdate(
		c.Request.Context(),
		&data,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "1")
}

// GetUserSellByID godoc
// @ID get_user_sell_by_id
// @Router /user/sell/{id} [GET]
// @Summary Get User Sell  By ID
// @Description Get User Sell  By ID
// @Tags User Sell
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.UserTransactionSell} "CoinBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetByIdTransactionSell(c *gin.Context) {

	Id := c.Param("id")
	if !utils.IsValidUUID(Id) {
		h.handleResponse(c, http.InvalidArgument, "User Sell id is an invalid uuid")
		return
	}

	resp, err := h.services.UserTransaction().GetByIdTransactionSell(
		context.Background(),
		&users_service.TransactioPrimaryKey{
			Id: Id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUserSellByID godoc
// @ID get_user_buy_by_id
// @Router /user/buy/{id} [GET]
// @Summary Get User Buy  By ID
// @Description Get User Buy  By ID
// @Tags User Buy
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.UserTransactionBuy} "UserTransactionBuyBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetByIdTransactionBuy(c *gin.Context) {

	Id := c.Param("id")
	if !utils.IsValidUUID(Id) {
		h.handleResponse(c, http.InvalidArgument, "User Buy id is an invalid uuid")
		return
	}

	resp, err := h.services.UserTransaction().GetByIdTransactionBuy(
		context.Background(),
		&users_service.TransactioPrimaryKey{
			Id: Id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, resp)
}

// GetHistoryTransactionUser godoc
// @ID get_user_history_transaction_user_by_id
// @Router /user/all-transfer/{id} [GET]
// @Summary Get User History Transaction User  By ID
// @Description Get User History Transaction User  By ID
// @Tags User Transfer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.HistoryUserTransaction} "HistoryUserTransactionBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetHistoryTransactionUser(c *gin.Context) {

	Id := c.Param("id")
	if !utils.IsValidUUID(Id) {
		h.handleResponse(c, http.InvalidArgument, "User Buy id is an invalid uuid")
		return
	}

	resp, err := h.services.UserTransaction().GetHistoryTransactionUser(
		context.Background(),
		&users_service.HistoryUserTransactionPrimaryKey{
			UserId: Id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, resp)
}

// GetTransactionSuccessImg godoc
// @ID get_transaction_success_img
// @Router /success-img/{id} [GET]
// @Summary Get Transaction Success Img By Id
// @Description Get Transaction Success Img By Id
// @Tags User Transfer Success Img
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.GetTransactionSuccessImgResponse} "GetTransactionSuccessImgResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTransactionSuccessImg(c *gin.Context) {

	Id := c.Param("id")
	if !utils.IsValidUUID(Id) {
		h.handleResponse(c, http.InvalidArgument, "Transaction success id is an invalid uuid")
		return
	}

	resp, err := h.services.UserTransaction().GetTransactionSuccessImg(
		context.Background(),
		&users_service.GetTransactionSuccessImgRequest{
			UserTransactionId: Id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, resp)
}
