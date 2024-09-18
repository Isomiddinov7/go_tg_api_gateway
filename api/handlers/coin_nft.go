package handlers

import (
	"context"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/coins_service"
	"go_tg_api_gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CoinNFTCreate godoc
// @ID create_coin_nft
// @Router /coin/nft [POST]
// @Summary Create CoinNFT
// @Description  Create CoinNFT
// @Tags CoinNFT
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Param nft_name formData string true "CoinNFT Name"
// @Param coin_nft_price formData string true "CoinNFT  Price"
// @Param nft_address formData string true "NFTAddress"
// @Success 200 {object} http.Response{data=coins_service.CoinNFT} "GetCoinNFTBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CoinNFTCreate(c *gin.Context) {

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

	var coin coins_service.CoinNFTCreate
	coin.NftName = c.PostForm("nft_name")
	coin.NftPrice = c.PostForm("coin_nft_price")
	coin.NftAddress = c.PostForm("nft_address")
	coin.NftImg = imageURL

	_, err = h.services.CoinNftService().Create(
		context.Background(),
		&coin,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.handleResponse(c, http.Created, gin.H{"message": "CoinNFT created successfully"})
}

// GetCoinNFTByID godoc
// @ID get_coin_nft_by_id
// @Router /coin/nft/{id} [GET]
// @Summary Get Coin  By ID
// @Description Get CoinNFT  By ID
// @Tags CoinNFT
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=coins_service.CoinNFT} "CoinNFTBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCoinNFTByID(c *gin.Context) {

	CoinID := c.Param("id")
	if !utils.IsValidUUID(CoinID) {
		h.handleResponse(c, http.InvalidArgument, "CoinNFT id is an invalid uuid")
		return
	}

	resp, err := h.services.CoinNftService().GetById(
		context.Background(),
		&coins_service.CoinNFTPrimaryKey{
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
// @ID get_coin_nft_list
// @Router /coin/nft [GET]
// @Summary Get CoinNFT List
// @Description  Get CoinNFT List
// @Tags CoinNFT
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=coins_service.GetListCoinNFTResponse} "GetListCoinNFTResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCoinNFTList(c *gin.Context) {

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

	resp, err := h.services.CoinNftService().GetList(
		c.Request.Context(),
		&coins_service.GetListCoinNFTRequest{
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

// UpdateCoinNFT godoc
// @ID update_coin_nft
// @Router /coin/nft/{id}  [PUT]
// @Summary Update CoinNFT
// @Description Update CoinNFt
// @Tags CoinNFT
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body coins_service.CoinNFTUpdate true "UpdateCoinNFT"
// @Success 200 {object} http.Response{data=coins_service.CoinNFT} "CoinNFT data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateCoinNFT(c *gin.Context) {

	var CoinNFT coins_service.CoinNFTUpdate
	CoinNFT.Id = c.Param("id")
	if !utils.IsValidUUID(CoinNFT.Id) {
		h.handleResponse(c, http.InvalidArgument, "CoinNFT id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&CoinNFT)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.CoinNftService().Update(
		c.Request.Context(),
		&CoinNFT,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteCoinNFT godoc
// @ID delete_coin_nft
// @Router /coin/nft/{id}  [DELETE]
// @Summary Delete CoinNFT
// @Description Delete CoinNFT
// @Tags CoinNFT
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "CoinNFT data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteCoinNFT(c *gin.Context) {

	CoinNFTId := c.Param("id")
	if !utils.IsValidUUID(CoinNFTId) {
		h.handleResponse(c, http.InvalidArgument, "CoinNFT id is an invalid uuid")
		return
	}

	_, err := h.services.CoinNftService().Delete(
		c.Request.Context(),
		&coins_service.CoinNFTPrimaryKey{Id: CoinNFTId},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "delete is success")
}
