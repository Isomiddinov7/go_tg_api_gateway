package handlers

import (
	"context"
	"encoding/json"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/coins_service"
	"go_tg_api_gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CreateCoin godoc
// @ID create_coin
// @Router /coin [POST]
// @Summary Create Coin
// @Description  Create Coin
// @Tags Coin
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Param name formData string true "Coin Name"
// @Param coin_buy_price formData string true "Coin Buy Price"
// @Param coin_sell_price formData string true "Coin Sell Price"
// @Param address formData string true "Address"
// @Param card_number formData string true "Card Number"
// @Param status formData string true "Status"
// @Param halfcoins formData array false "coins_service.CreateCoin.HalfCoinPrice" items.schema.type=object items.schema.$ref="#/definitions/HalfCoinPrice"
// @Success 200 {object} http.Response{data=coins_service.Coin} "GetCoinBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateCoin(c *gin.Context) {

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

	var coin coins_service.CreateCoin

	coin.Name = c.PostForm("name")
	coin.CoinBuyPrice = c.PostForm("coin_buy_price")
	coin.CoinSellPrice = c.PostForm("coin_sell_price")
	coin.Address = c.PostForm("address")
	coin.CardNumber = c.PostForm("card_number")
	coin.Status = c.PostForm("status")
	coin.ImageId = imageURL
	halfcoinsJSON := c.PostForm("halfcoins")
	var halfcoins []*coins_service.HalfCoinPrice
	if err := json.Unmarshal([]byte(halfcoinsJSON), &halfcoins); err != nil {
		h.handleResponse(c, http.BadRequest, gin.H{"error": "Invalid halfcoins data"})
		return
	}

	coin.Halfcoins = halfcoins
	_, err = h.services.CoinService().Create(
		c.Request.Context(),
		&coin,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.handleResponse(c, http.Created, gin.H{"message": "Coin created successfully"})
}

// GetCoinByID godoc
// @ID get_coin_by_id
// @Router /coin/{id} [GET]
// @Summary Get Coin  By ID
// @Description Get Coin  By ID
// @Tags Coin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=coins_service.Coin} "CoinBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCoinByID(c *gin.Context) {

	CoinID := c.Param("id")
	if !utils.IsValidUUID(CoinID) {
		h.handleResponse(c, http.InvalidArgument, "Coin id is an invalid uuid")
		return
	}

	resp, err := h.services.CoinService().GetById(
		context.Background(),
		&coins_service.CoinPrimaryKey{
			Id: CoinID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetCoinList godoc
// @ID get_coin_list
// @Router /coin [GET]
// @Summary Get Coins List
// @Description  Get Coins List
// @Tags Coin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=coins_service.GetListCoinResponse} "GetAllCoinResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCoinList(c *gin.Context) {

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

	resp, err := h.services.CoinService().GetList(
		c.Request.Context(),
		&coins_service.GetListCoinRequest{
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

// UpdateCoin godoc
// @ID update_coin
// @Router /v1/coin/{id} [PUT]
// @Summary Update Coin
// @Description Update Coin
// @Tags Coin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body coins_service.UpdateCoin true "UpdateCoin"
// @Success 200 {object} http.Response{data=coins_service.Coin} "Coin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateCoin(c *gin.Context) {

	var Coin coins_service.UpdateCoin
	Coin.Id = c.Param("id")
	if !utils.IsValidUUID(Coin.Id) {
		h.handleResponse(c, http.InvalidArgument, "Coin id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Coin)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.CoinService().Update(
		c.Request.Context(),
		&Coin,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteCoin godoc
// @ID delete_coin
// @Router /coin/{id} [DELETE]
// @Summary Delete Coin
// @Description Delete Coin
// @Tags Coin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Coin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteCoin(c *gin.Context) {

	CoinId := c.Param("id")
	if !utils.IsValidUUID(CoinId) {
		h.handleResponse(c, http.InvalidArgument, "Coin id is an invalid uuid")
		return
	}

	_, err := h.services.CoinService().Delete(
		c.Request.Context(),
		&coins_service.CoinPrimaryKey{Id: CoinId},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "delete is success")
}
