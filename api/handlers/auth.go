package handlers

import (
	"errors"
	"fmt"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/config"
	"go_tg_api_gateway/genproto/users_service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary login
// @Description login
// @Tags Admin
// @Accept json
// @Produce json
// @Param object body users_service.Credentials true "LoginRequestBody"
// @Success 200 {object} http.Response{data=string} "Token"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Login(c *gin.Context) {
	var credentials users_service.Credentials

	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	resp, err := h.services.Auth().Login(
		c.Request.Context(),
		&credentials,
	)
	if err != nil {
		h.handleResponse(c, http.BadRequest, gin.H{"error": "login or password is wrong"})
		return
	}
	fmt.Println(resp)
	token, err := GenerateToken(resp.Id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	h.handleResponse(c, http.OK, gin.H{"token": token})
}

func GenerateToken(id string) (string, error) {
	expirationTime := time.Now().Add(config.ExpiredTime)
	claims := &jwt.StandardClaims{
		Id:        id,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	fmt.Println(tokenString, err)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenStr string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid signature")
		}
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if int64(claims.ExpiresAt) < time.Now().Unix() {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
