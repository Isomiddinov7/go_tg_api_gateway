package handlers

import (
	"context"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/coins_service"
	"go_tg_api_gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CreateNFT godoc
// @ID create_nft
// @Router /nft [POST]
// @Summary Create NFT
// @Description  Create NFT
// @Tags NFT
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Param comment formData string true "NFT Comment"
// @Param user_id formData string true "User ID NFT"
// @Success 200 {object} http.Response{data=coins_service.Coin} "GetCoinBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateNFT(c *gin.Context) {

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

	var nft coins_service.CreateNFT

	nft.Comment = c.PostForm("comment")
	nft.UserId = c.PostForm("user_id")
	nft.NftImg = imageURL

	_, err = h.services.NFT().Create(
		c.Request.Context(),
		&nft,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.handleResponse(c, http.Created, gin.H{"message": "NFT created successfully"})
}

// GetNFTByID godoc
// @ID get_nft_by_id
// @Router /nft/{id} [GET]
// @Summary Get NFT  By ID
// @Description Get NFT  By ID
// @Tags NFT
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=coins_service.NFT} "NFTBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetNFTByID(c *gin.Context) {

	nftID := c.Param("id")
	if !utils.IsValidUUID(nftID) {
		h.handleResponse(c, http.InvalidArgument, "NFT id is an invalid uuid")
		return
	}

	resp, err := h.services.NFT().GetById(
		context.Background(),
		&coins_service.NFTPrimaryKey{
			Id: nftID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetNFTList godoc
// @ID get_nft_list
// @Router /nft [GET]
// @Summary Get NFTs List
// @Description  Get NFTs List
// @Tags NFT
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=coins_service.GetListNFTResponse} "GetListNFTResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetNFTList(c *gin.Context) {

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

	resp, err := h.services.NFT().GetAll(
		c.Request.Context(),
		&coins_service.GetListNFTRequest{
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

// UpdateNFT godoc
// @ID update_nft
// @Router /nft/{id} [PUT]
// @Summary Update NFT
// @Description Update NFT
// @Tags NFT
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body coins_service.UpdateNFT true "UpdateNFT"
// @Success 200 {object} http.Response{data=coins_service.NFT} "NFT data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateNFT(c *gin.Context) {

	var NFT coins_service.UpdateNFT
	NFT.Id = c.Param("id")
	if !utils.IsValidUUID(NFT.Id) {
		h.handleResponse(c, http.InvalidArgument, "NFT id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&NFT)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.NFT().Update(
		c.Request.Context(),
		&NFT,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteNFT godoc
// @ID delete_nft
// @Router /nft/{id} [DELETE]
// @Summary Delete NFT
// @Description Delete NFT
// @Tags NFT
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "NFT data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteNFT(c *gin.Context) {

	NFTId := c.Param("id")
	if !utils.IsValidUUID(NFTId) {
		h.handleResponse(c, http.InvalidArgument, "NFT id is an invalid uuid")
		return
	}

	_, err := h.services.NFT().Delete(
		c.Request.Context(),
		&coins_service.NFTPrimaryKey{Id: NFTId},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "delete is success")
}
