package main

import (
	"fmt"
	"go_tg_api_gateway/api"
	"go_tg_api_gateway/api/handlers"
	"go_tg_api_gateway/config"
	"go_tg_api_gateway/pkg/logger"
	"go_tg_api_gateway/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	grpcSvcs, err := services.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	var loggerLevel = new(string)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("go_tg_api_gateway", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, log, grpcSvcs)

	api.SetUpAPI(r, h, cfg)

	fmt.Println("Start api gateway....")

	if err := r.Run(cfg.HTTPPort); err != nil {
		return
	}
}
