package handlers

import (
	"context"
	"fmt"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/users_service"
	"go_tg_api_gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CreateMessage godoc
// @ID create_user_message
// @Router /user/message [POST]
// @Summary Create Message
// @Description  Create Message
// @Tags User
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "Upload file"
// @Param text formData string true "TEXT"
// @Param user_id formData string true "UserId"
// @Success 200 {object} http.Response{data=string} "GetCoinBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUserMessage(c *gin.Context) {

	var message users_service.MessageRequest
	file, err := c.FormFile("file")
	fmt.Println(file)
	if file != nil {
		if err != nil {
			h.handleResponse(c, http.ErrMissingFile, gin.H{"error": "http: no such file"})
		} else {
			imageURL, err := utils.UploadImage(file)
			if err != nil {
				h.handleResponse(c, http.InternalServerError, gin.H{"error": "Failed to upload image"})
				return
			}
			message.File = imageURL

			message.Text = c.PostForm("text")
			message.UserId = c.PostForm("user_id")

			_, err = h.services.Messages().CreateUserMessage(
				c.Request.Context(),
				&message,
			)
			if err != nil {
				h.handleResponse(c, http.InternalServerError, gin.H{"error": err.Error()})
				return
			}

			h.handleResponse(c, http.Created, gin.H{"message": "Message created successfully"})

		}
	} else {

		message.File = ""
		message.Text = c.PostForm("text")
		message.UserId = c.PostForm("user_id")

		_, err = h.services.Messages().CreateUserMessage(
			c.Request.Context(),
			&message,
		)
		if err != nil {
			h.handleResponse(c, http.InternalServerError, gin.H{"error": err.Error()})
			return
		}

		h.handleResponse(c, http.Created, gin.H{"message": "Message created successfully"})
	}

}

// CreateMessage godoc
// @ID create_admin_message
// @Router /admin/message [POST]
// @Summary Create Message
// @Description  Create Message
// @Tags Admin
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "Upload file"
// @Param text formData string true "TEXT"
// @Param user_id formData string true "UserId"
// @Success 200 {object} http.Response{data=string} "GetCoinBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdminMessage(c *gin.Context) {

	var message users_service.MessageRequest
	file, err := c.FormFile("file")
	if file != nil {
		if err != nil {
			h.handleResponse(c, http.ErrMissingFile, gin.H{"error": "http: no such file"})
		} else {
			imageURL, err := utils.UploadImage(file)
			if err != nil {
				h.handleResponse(c, http.InternalServerError, gin.H{"error": "Failed to upload image"})
				return
			}
			message.File = imageURL
		}
	} else {
		message.File = ""
		message.Text = c.PostForm("text")
		message.UserId = c.PostForm("user_id")
		_, err = h.services.Messages().CreateAdminMessage(
			c.Request.Context(),
			&message,
		)
		if err != nil {
			h.handleResponse(c, http.InternalServerError, gin.H{"error": err.Error()})
			return
		}

		h.handleResponse(c, http.Created, gin.H{"message": "Message created successfully"})
	}
}

// UpdateMessage godoc
// @ID update_message
// @Router /message/{id} [PUT]
// @Summary Update Message
// @Description Update Message
// @Tags Message
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateMessageRequest true "UpdateMessage"
// @Success 200 {object} http.Response{data=string} "Message data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateMessage(c *gin.Context) {

	var message users_service.UpdateMessageRequest
	message.Id = c.Param("id")
	if !utils.IsValidUUID(message.Id) {
		h.handleResponse(c, http.InvalidArgument, "Message id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&message)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.Messages().UpdateMessage(
		c.Request.Context(),
		&message,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUserMessageByID godoc
// @ID get_user_message_by_id
// @Router /user/message/{id} [GET]
// @Summary Get User Message  By ID
// @Description Get User Message  By ID
// @Tags User Message
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.GetMessageUserResponse} "UserMessageBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserMessage(c *gin.Context) {
	user_id := c.Param("id")
	if !utils.IsValidUUID(user_id) {
		h.handleResponse(c, http.InvalidArgument, "User id is an invalid uuid")
		return
	}

	resp, err := h.services.Messages().GetUserMessage(
		context.Background(),
		&users_service.GetMessageUserRequest{
			UserId: user_id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

// GetUserMessageByID godoc
// @ID get_admin_all_message_by_id
// @Router /admin/message/ [GET]
// @Summary Get Admin Message  By ID
// @Description Get Admin Message  By ID
// @Tags Admin Message
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=users_service.GetMessageAdminResponse} "UserMessageBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminAllMessage(c *gin.Context) {

	resp, err := h.services.Messages().GetAdminAllMessage(
		context.Background(),
		&users_service.Empty{},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

// GetAdminMessageByID godoc
// @ID get_admin_message_by_id
// @Router /admin/message/{id} [GET]
// @Summary Get Admin Message  By ID
// @Description Get Admin Message  By ID
// @Tags Admin Message
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.GetMessageUserResponse} "UserMessageBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminMessage(c *gin.Context) {
	user_id := c.Param("id")
	if !utils.IsValidUUID(user_id) {
		h.handleResponse(c, http.InvalidArgument, "User id is an invalid uuid")
		return
	}

	resp, err := h.services.Messages().GetUserMessage(
		context.Background(),
		&users_service.GetMessageUserRequest{
			UserId: user_id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

// GetMessageAdminID godoc
// @ID get_message_admin_id
// @Router /admin/message/user/{id} [GET]
// @Summary Get Admin Message  By ID
// @Description Get Admin Message  By ID
// @Tags Admin Message
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.GetMessageUserResponse} "UserMessageBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetMessageAdminID(c *gin.Context) {
	user_id := c.Param("id")
	if !utils.IsValidUUID(user_id) {
		h.handleResponse(c, http.InvalidArgument, "User id is an invalid uuid")
		return
	}

	resp, err := h.services.Messages().GetMessageAdminID(
		context.Background(),
		&users_service.GetMessageUserRequest{
			UserId: user_id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
