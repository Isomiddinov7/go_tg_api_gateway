package models

import "github.com/golang-jwt/jwt"



type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}
