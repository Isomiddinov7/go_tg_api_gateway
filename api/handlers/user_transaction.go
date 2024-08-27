package handlers

import (
	"context"
	"fmt"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/coins_service"
	"go_tg_api_gateway/genproto/users_service"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserSell godoc
// @ID User_sell
// @Router /v1/user/sell [POST]
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

	ext := filepath.Ext(file.Filename)
	validFormats := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	if !validFormats[ext] {
		h.handleResponse(c, http.BadRequest, gin.H{"error": "Invalid file format"})
		return
	}

	err = os.MkdirAll("images/user_sell", os.ModePerm)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Unable to create uploads directory"})
		return
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1) + ext
	filePath := filepath.Join("images/user_sell", filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	imageLink := filePath
	imgId, err := h.services.FileImage().ImageUpload(
		context.Background(),
		&coins_service.ImageData{
			Id:        uniqueId.String(),
			ImageLink: imageLink,
		},
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Unable to save file info to database"})
		return
	}

	var user_sell users_service.UserSellRequest

	user_sell.UserId = c.PostForm("userId")
	user_sell.CoinId = c.PostForm("coinId")
	user_sell.CoinAmount = c.PostForm("coin_amount")
	user_sell.Message = c.PostForm("message")
	user_sell.CardNumber = c.PostForm("card_number")
	user_sell.CheckImg = imgId.Id

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
// @Router /v1/user/buy [POST]
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

	ext := filepath.Ext(file.Filename)
	validFormats := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	if !validFormats[ext] {
		h.handleResponse(c, http.BadRequest, gin.H{"error": "Invalid file format"})
		return
	}

	err = os.MkdirAll("images/user_buy", os.ModePerm)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Unable to create uploads directory"})
		return
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1) + ext
	filePath := filepath.Join("images/user_buy", filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	imageLink := filePath
	imgId, err := h.services.FileImage().ImageUpload(
		context.Background(),
		&coins_service.ImageData{
			Id:        uniqueId.String(),
			ImageLink: imageLink,
		},
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Unable to save file info to database"})
		return
	}

	var user_buy users_service.UserBuyRequest
	user_buy.UserId = c.PostForm("userId")
	user_buy.CoinId = c.PostForm("coinId")
	user_buy.CoinAmount = c.PostForm("coin_amount")
	user_buy.Message = c.PostForm("message")
	user_buy.Address = c.PostForm("address")
	user_buy.PayImg = imgId.Id

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
// @Router /v1/user/sell [GET]
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

	for i := range resp.UserTransaction {
		fileName := strings.Replace(resp.UserTransaction[i].CheckImg, "-", "", -1)
		extensions := []string{".png", ".gif", ".jpg", ".jpeg"}

		var filePath, imageUrl string

		var found bool

		for _, ext := range extensions {
			potensialPath := fmt.Sprintf("./images/user_sell/%s%s", fileName, ext)
			link := fmt.Sprintf("user/image/%s%s", fileName, ext)
			if _, err := os.Stat(potensialPath); err == nil {
				filePath = potensialPath
				imageUrl = link
				found = true
				break
			}
		}

		if !found {
			h.handleResponse(c, http.NoContent, gin.H{"error": "File type not found"})
			return
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			h.handleResponse(c, http.NoContent, gin.H{"error": "File not found"})
			return
		}

		fileURL := fmt.Sprintf("https://alimkulov.uz/%s", imageUrl)
		resp.UserTransaction[i].CheckImg = fileURL
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// AllUserBuy godoc
// @ID get_transaction_by_list
// @Router /v1/user/buy [GET]
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

	for i := range resp.UserTransaction {
		fileName := strings.Replace(resp.UserTransaction[i].CheckImg, "-", "", -1)
		extensions := []string{".png", ".gif", ".jpg", ".jpeg"}

		var filePath, imageUrl string

		var found bool

		for _, ext := range extensions {
			potensialPath := fmt.Sprintf("./images/user_buy/%s%s", fileName, ext)
			link := fmt.Sprintf("user/image/%s%s", fileName, ext)
			if _, err := os.Stat(potensialPath); err == nil {
				filePath = potensialPath
				imageUrl = link
				found = true
				break
			}
		}

		if !found {
			h.handleResponse(c, http.NoContent, gin.H{"error": "File type not found"})
			return
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			h.handleResponse(c, http.NoContent, gin.H{"error": "File not found"})
			return
		}

		fileURL := fmt.Sprintf("https://alimkulov.uz/%s", imageUrl)
		resp.UserTransaction[i].CheckImg = fileURL
	}

	h.handleResponse(c, http.OK, resp)
}
