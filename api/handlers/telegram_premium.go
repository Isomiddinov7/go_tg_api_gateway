package handlers

import (
	"context"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/coins_service"
	"go_tg_api_gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CreatePrice godoc
// @ID create_premium_price
// @Router /premium/price [POST]
// @Summary Create Price
// @Description  Create Price
// @Tags Premium
// @Accept json
// @Produce json
// @Param profile body coins_service.CreateTelegramPremiumPrice true "CreateTelegramPremiumPriceBody"
// @Success 200 {object} http.Response{data=coins_service.TelegramPremiumPrice} "GetTelegramPremiumPriceBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreatePrice(c *gin.Context) {

	var price coins_service.CreateTelegramPremiumPrice

	err := c.ShouldBindJSON(&price)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.services.TelegramPremium().CreatePrice(
		c.Request.Context(),
		&price,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "ok")
}

// CreatePremium godoc
// @ID create_premium
// @Router /premium/create [POST]
// @Summary Create Premium
// @Description  Create Premium
// @Tags Premium
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Param name formData string true "Premium Name"
// @Param card_number formData string true "Card Number"
// @Success 200 {object} http.Response{data=coins_service.TelegramPremium} "GetTelegramPremiumBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreatePremium(c *gin.Context) {

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

	var premium coins_service.CreateTelegramPremium

	premium.Name = c.PostForm("name")
	premium.CardNumber = c.PostForm("card_number")
	premium.Img = imageURL

	_, err = h.services.TelegramPremium().CreatePremium(
		c.Request.Context(),
		&premium,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.handleResponse(c, http.Created, gin.H{"message": "Coin created successfully"})
}

// GetPremiumById godoc
// @ID get_premium_by_id
// @Router /premium/{id} [GET]
// @Summary Get Premium  By ID
// @Description Get Premium  By ID
// @Tags Premium
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=coins_service.TelegramPremium} "TelegramPremiumBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPremiumById(c *gin.Context) {

	premiumID := c.Param("id")
	if !utils.IsValidUUID(premiumID) {
		h.handleResponse(c, http.InvalidArgument, "premium id is an invalid uuid")
		return
	}

	resp, err := h.services.TelegramPremium().GetPremiumById(
		context.Background(),
		&coins_service.TelegramPriemiumPrimaryKey{
			Id: premiumID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateTransactionStatus godoc
// @ID update_status
// @Router /premium/{id} [PUT]
// @Summary Update Premium
// @Description Update Premium
// @Tags Premium
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body coins_service.UpdateStatus true "UpdateStatusCoin"
// @Success 200 {object} http.Response{data=coins_service.TelegramPremium} "TelegramPremium data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateTransactionStatus(c *gin.Context) {

	var status coins_service.UpdateStatus
	status.TransactionId = c.Param("id")
	if !utils.IsValidUUID(status.TransactionId) {
		h.handleResponse(c, http.InvalidArgument, "Coin id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&status)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.TelegramPremium().UpdateTransactionStatus(
		c.Request.Context(),
		&status,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// CreatePremiumTransaction godoc
// @ID create_premium_transaction
// @Router /premium/transaction [POST]
// @Summary Create Premium Transaction
// @Description  Create Premium Transaction
// @Tags Premium
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Param phone_number formData string true "Phone number"
// @Param telegram_username formData string true "Telegram Username"
// @Param price_id formData string true "Price ID"
// @Param user_id formData string true "User ID"
// @Success 200 {object} http.Response{data=coins_service.TelegramPremium} "GetTelegramPremiumBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) PremiumTransaction(c *gin.Context) {

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

	var premium coins_service.PremiumTransactionRequest

	premium.PhoneNumber = c.PostForm("phone_number")
	premium.TelegramUsername = c.PostForm("telegram_username")
	premium.TelegramPriceId = c.PostForm("price_id")
	premium.UserId = c.PostForm("user_id")
	premium.PaymentImg = imageURL

	_, err = h.services.TelegramPremium().PremiumTransaction(
		c.Request.Context(),
		&premium,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.handleResponse(c, http.Created, gin.H{"message": "Coin created successfully"})
}

// @Security ApiKeyAuth
// GetPremiumTransactionList godoc
// @ID get_premium_transaction_list
// @Router /premium/transaction [GET]
// @Summary Get Premium Transaction List
// @Description  Get Premium Transaction List
// @Tags Premium
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=coins_service.GetPremiumTransactionResponse} "GetPremiumTransactionResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetList(c *gin.Context) {

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

	resp, err := h.services.TelegramPremium().GetList(
		c.Request.Context(),
		&coins_service.GetListPremiumRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetPremiumList godoc
// @ID get_premium_list
// @Router /premium [GET]
// @Summary Get Premium Transaction List
// @Description  Get Premium Transaction List
// @Tags Premium
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=coins_service.GetPremiumListResponse} "GetPremiumListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPremiumList(c *gin.Context) {

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

	resp, err := h.services.TelegramPremium().GetPremiumList(
		c.Request.Context(),
		&coins_service.GetPremiumListRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, resp)
}

// GetPremiumTransactionByID godoc
// @ID get_premium_transaction_id
// @Router /premium/transaction/{id} [GET]
// @Summary Get PremiumTransaction  By ID
// @Description Get PremiumTransaction  By ID
// @Tags Premium
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=coins_service.GetPremiumTransactionId} "GetPremiumTransactionIdBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPremiumTransactionById(c *gin.Context) {

	premium_transaction_id := c.Param("id")
	if !utils.IsValidUUID(premium_transaction_id) {
		h.handleResponse(c, http.InvalidArgument, "premium_transaction_id id is an invalid uuid")
		return
	}

	resp, err := h.services.TelegramPremium().GetPremiumTransactionById(
		context.Background(),
		&coins_service.GetPremiumTransactionPrimaryKey{
			PremiumTransactionId: premium_transaction_id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
