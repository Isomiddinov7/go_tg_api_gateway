package config

import "time"

const (
	CtxTimeout = time.Second * 2

	ExpiredTime = time.Hour * 24
)

var JwtKey = []byte("shHQjqnDKN")
