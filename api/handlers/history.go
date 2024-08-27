package handlers

import (
	"context"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/coins_service"
	"go_tg_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// GetHistoryUserList godoc
// @ID get_history_user_list
// @Router /v1/history/user [GET]
// @Summary Get History Users List
// @Description  Get History Users List
// @Tags History
// @Accept json
// @Produce json
// @Param user_id query integer true "user_id"
// @Success 200 {object} http.Response{data=coins_service.HistoryUserResponse} "GetHistoryUserResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) HistoryUser(c *gin.Context) {

	UserID := c.Param("id")

	if !util.IsValidUUID(UserID) {
		h.handleResponse(c, http.InvalidArgument, "User id is an invalid uuid")
		return
	}

	resp, err := h.services.History().HistoryUser(
		context.Background(),
		&coins_service.HistoryUserRequest{UserId: UserID},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}