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
	r.POST("/login", h.Login)

	v1 := r.Group("v1")
	v1.Use(h.AuthMiddleware())
	{
		v1.POST("/coin", h.CreateCoin)
		v1.GET("/coin/:id", h.GetCoinByID)
		v1.GET("/coin", h.GetCoinList)
		v1.PUT("/coin/:id", h.UpdateCoin)
		v1.DELETE("/coin/:id", h.DeleteCoin)

		//Buy Or Sale
		v1.POST("/sell", h.GetSell)
		v1.POST("/buy", h.GetBuy)

		// USER
		v1.POST("/user", h.CreateUser)
		v1.GET("/user/:id", h.GetUserByID)
		v1.GET("/user", h.GetUserList)
		v1.PUT("/user/:id", h.UpdateUser)

		//Transactions
		v1.POST("/user/sell", h.UserSell)
		v1.POST("/user/buy", h.UserBuy)
		v1.GET("/user/sell", h.AllUserSell)
		v1.GET("/user/buy", h.AllUserBuy)

		//Messages
		v1.POST("/user/message", h.CreateUserMessage)
		v1.POST("/admin/message", h.CreateAdminMessage)
		v1.PUT("/message/:id", h.UpdateMessage)
		v1.GET("/user/message/:id", h.GetUserMessage)
		v1.GET("/admin/message", h.GetAdminAllMessage)
		v1.GET("/admin/message/:id", h.GetUserMessage)
		v1.GET("/admin/message/user/:id", h.GetMessageAdminID)

		//History
		v1.GET("/history/user", h.HistoryUser)

	}

	// COIN

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
