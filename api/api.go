package api

import (
	"go_tg_api_gateway/api/handlers"
	"go_tg_api_gateway/config"
	"go_tg_api_gateway/middlewares"

	_ "go_tg_api_gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	r.Use(middlewares.Cors())
	r.Use(MaxAllowed(5000))

	// r.Use(h.CheckPasswordMiddleware())
	r.POST("/login", h.Login)

	v1 := r.Group("v1")
	// v1.Use(h.AuthMiddleware())
	{
		v1.POST("/coin", h.CreateCoin)
		v1.GET("/coin/:id", h.GetCoinByID)
		v1.GET("/coin", h.GetCoinList)
		v1.PUT("/coin/:id", h.UpdateCoin)
		v1.DELETE("/coin/:id", h.DeleteCoin)

		//Buy Or Sale
		v1.POST("/sell", h.GetSell)
		v1.POST("/buy", h.GetBuy)

		//Transactions
		v1.POST("/user/sell", h.UserSell)
		v1.POST("/user/buy", h.UserBuy)
		v1.GET("/user/sell", h.AllUserSell)
		v1.GET("/user/buy", h.AllUserBuy)

		//Messages
		v1.POST("/admin/message", h.CreateAdminMessage)
		v1.PUT("/message/:id", h.UpdateMessage)
		v1.GET("/admin/message", h.GetAdminAllMessage)
		v1.GET("/admin/message/:id", h.GetUserMessage)
		v1.GET("/admin/message/user/:id", h.GetMessageAdminID)

		//History
		v1.GET("/history/user", h.HistoryUser)

	}

	// USER
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUserByID)
	r.GET("/user", h.GetUserList)
	r.PUT("/user/:id", h.UpdateUser)
	r.POST("/user/message", h.CreateUserMessage)
	r.GET("/user/message/:id", h.GetUserMessage)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
