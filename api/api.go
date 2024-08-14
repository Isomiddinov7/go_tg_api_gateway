package api

import (
	"go_tg_api_gateway/api/handlers"
	"go_tg_api_gateway/config"

	_ "go_tg_api_gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	r.Use(customCORSMiddleware())
	r.Use(MaxAllowed(5000))

	// r.Use(h.CheckPasswordMiddleware())

	// r.POST("/login", h.Login)

	// COIN
	r.POST("/coin", h.CreateCoin)
	r.GET("/coin/:id", h.GetCoinByID)
	r.GET("/coin", h.GetCoinList)
	r.PUT("/coin/:id", h.UpdateCoin)
	r.DELETE("/coin/:id", h.DeleteCoin)
	// r.POST("/coin/image", h.HandleCoinFileupload)
	// r.GET("/coin/image/:id", h.GetCoinFileURL)
	// r.DELETE("/coin/image/:id", h.HandleCoinDeleteImage)

	//Buy Or Sale
	r.POST("/sell", h.GetSell)
	r.POST("/user/sell", h.UserSell)
	r.POST("/user/image/sell", h.HandleUserSellFileupload)
	r.GET("/user/image/sell/:id", h.GetUserSellFileURL)
	r.DELETE("/user/image/sell/:id", h.HandleUserSellDeleteImage)
	r.POST("/buy", h.GetBuy)
	r.POST("/user/buy", h.UserBuy)
	r.POST("/user/image/buy", h.HandleUserBuyFileupload)
	r.GET("/user/image/buy/:id", h.GetUserBuyFileURL)
	r.DELETE("/user/image/buy/:id", h.HandleUserBuyDeleteImage)

	// USER
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUserByID)
	r.GET("/user", h.GetUserList)
	r.PUT("/user/:id", h.UpdateUser)
	// r.POST("/user/image", h.HandleUserFileupload)
	// r.GET("/user/image/:id", h.GetUserFileURL)
	// r.DELETE("/user/image/:id", h.HandleUserDeleteImage)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func MaxAllowed(n int) gin.HandlerFunc {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
	}

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request

		c.Set("sem", sem)
		c.Set("count_request", countReq)

		c.Next()
	}
}
