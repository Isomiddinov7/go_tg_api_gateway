package handlers

import (
	"context"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/users_service"
	"go_tg_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @ID create_user
// @Router /v1/user [POST]
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
// @Router /v1/user/{id} [GET]
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

	UserID := c.Param("id")

	if !util.IsValidUUID(UserID) {
		h.handleResponse(c, http.InvalidArgument, "User id is an invalid uuid")
		return
	}

	resp, err := h.services.UserService().GetById(
		context.Background(),
		&users_service.UserPrimaryKey{
			Id: UserID,
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
// @Router /v1/user [GET]
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

// UpdateUser godoc
// @ID update_user
// @Router /v1/user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateUser true "UpdateUser"
// @Success 200 {object} http.Response{data=users_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {

	var user users_service.UpdateUser

	user.Id = c.Param("id")

	if !util.IsValidUUID(user.Id) {
		h.handleResponse(c, http.InvalidArgument, "User id is an invalid uuid")
		return
	}

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

	h.handleResponse(c, http.OK, "1")
}
