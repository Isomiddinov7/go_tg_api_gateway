package handlers

import (
	"context"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/users_service"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description  Create User
// @Tags User
// @Accept json
// @Produce json
// @Param profile body users_service.CreateUser true "CreateUserBody"
// @Success 200 {object} http.Response{data=users_service.User} "GetUserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {

	var user users_service.CreateUser

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.services.UserService().Create(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "ok")
}

// GetUserByID godoc
// @ID get_user_by_id
// @Router /user/{id} [GET]
// @Summary Get User  By ID
// @Description Get User  By ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.User} "UserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserByID(c *gin.Context) {

	TelegramId := c.Param("id")
	resp, err := h.services.UserService().GetById(
		context.Background(),
		&users_service.UserPrimaryKey{
			Id: TelegramId,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetUserList godoc
// @ID get_user_list
// @Router /user [GET]
// @Summary Get Users List
// @Description  Get Users List
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=users_service.GetListUserResponse} "GetAllUserResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserList(c *gin.Context) {

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

	resp, err := h.services.UserService().GetList(
		context.Background(),
		&users_service.GetListUserRequest{
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
// UpdateUserStatus godoc
// @ID update_user_status
// @Router /user [PUT]
// @Summary Update User Status
// @Description  Update User Status
// @Tags User
// @Accept json
// @Produce json
// @Param profile body users_service.UpdateUserStatus true "UpdateUserStatusBody"
// @Success 200 {object} http.Response{data=string} "Update User Status Body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Update(c *gin.Context) {
	var user users_service.UpdateUserStatus

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.services.UserService().Update(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "ok")
}
