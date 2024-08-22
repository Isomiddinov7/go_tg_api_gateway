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
	r.POST("/coin-service/coin", h.CreateCoin)
	r.GET("/coin-service/coin/:id", h.GetCoinByID)
	r.GET("/coin-service/coin", h.GetCoinList)
	r.PUT("/coin-service/coin/:id", h.UpdateCoin)
	r.DELETE("/coin-service/coin/:id", h.DeleteCoin)

	//Buy Or Sale
	r.POST("/coin-service/sell", h.GetSell)
	r.POST("/coin-service/buy", h.GetBuy)

	// USER
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUserByID)
	r.GET("/user", h.GetUserList)
	r.PUT("/user/:id", h.UpdateUser)

	//Transactions
	r.POST("/user/sell", h.UserSell)
	r.POST("/user/buy", h.UserBuy)
	r.GET("/user/sell", h.AllUserSell)
	r.GET("/user/buy", h.AllUserBuy)

	//Messages
	r.POST("/user/message", h.CreateUserMessage)
	r.POST("/admin/message", h.CreateAdminMessage)
	r.PUT("/message/:id", h.UpdateMessage)
	r.GET("/user/message/:id", h.GetUserMessage)
	r.GET("/admin/message", h.GetAdminAllMessage)
	r.GET("/admin/message/:id", h.GetUserMessage)

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
