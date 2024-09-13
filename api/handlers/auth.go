package handlers

import (
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/genproto/users_service"
	initializers "go_tg_api_gateway/initializer"
	"go_tg_api_gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary login
// @Description login
// @Tags Admin
// @Accept json
// @Produce json
// @Param object body users_service.Req true "LoginRequestBody"
// @Success 200 {object} http.Response{data=string} "Token"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Auth(c *gin.Context) {
	var credentials users_service.Req

	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	resp, err := h.services.Auth().Auth(
		c.Request.Context(),
		&credentials,
	)
	if err != nil {
		h.handleResponse(c, http.BadRequest, gin.H{"error": "login or password is wrong"})
		return
	}
	config, _ := initializers.LoadConfig(".")
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, resp.UserId, config.AccessTokenPrivateKey)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, resp.UserId, config.RefreshTokenPrivateKey)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	h.handleResponse(c, http.OK, gin.H{"token": access_token})
}
