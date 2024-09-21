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

	r.Use(middlewares.Cors())
	r.Use(MaxAllowed(5000))

	r.POST("/login", h.Auth)

	r.POST("/coin", h.DeserializeUser(), h.CreateCoin)
	// r.GET("/coin/:id", h.DeserializeUser(), h.GetCoinByID)
	// r.GET("/coin", h.DeserializeUser(), h.GetCoinList)
	r.PUT("/coin/:id", h.DeserializeUser(), h.UpdateCoin)
	r.DELETE("/coin/:id", h.DeserializeUser(), h.DeleteCoin)

	r.GET("/coin/:id", h.GetCoinByID)
	r.GET("/coin", h.GetCoinList)

	r.POST("/sell", h.GetSell)
	r.POST("/buy", h.GetBuy)

	r.POST("/user/sell", h.UserSell)
	r.POST("/user/buy", h.UserBuy)
	r.GET("/user/sell", h.AllUserSell)
	r.GET("/user/buy/:id", h.GetByIdTransactionBuy)
	r.GET("/user/sell/:id", h.GetByIdTransactionSell)
	r.GET("/user/buy", h.AllUserBuy)

	r.POST("/admin/message", h.DeserializeUser(), h.CreateAdminMessage)
	r.PUT("/message/:id", h.UpdateMessage)
	r.GET("/admin/message", h.DeserializeUser(), h.GetAdminAllMessage)
	// r.GET("/admin/message", h.GetAdminAllMessage)
	r.GET("/admin/message/:id", h.DeserializeUser(), h.GetUserMessage)
	r.GET("/admin/message/user/:id", h.DeserializeUser(), h.GetMessageAdminID)

	// r.POST("/admin/message", h.CreateAdminMessage)
	// r.PUT("/message/:id", h.UpdateMessage)
	// r.GET("/admin/message", h.GetAdminAllMessage)
	// r.GET("/admin/message/:id", h.GetUserMessage)
	// r.GET("/admin/message/user/:id", h.GetMessageAdminID)

	r.GET("/history/user", h.HistoryUser)

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUserByID)
	r.GET("/user", h.DeserializeUser(), h.GetUserList)
	r.PUT("/user/:id", h.UpdateUser)
	r.POST("/user/message", h.CreateUserMessage)
	r.GET("/user/message/:id", h.GetUserMessage)

	r.POST("/premium/price", h.DeserializeUser(), h.CreatePrice)
	r.POST("/premium/create", h.DeserializeUser(), h.CreatePremium)
	r.GET("/premium/:id", h.GetPremiumById)
	r.GET("/premium/transaction", h.GetList)
	r.GET("/premium/transaction/:id", h.DeserializeUser(), h.GetPremiumTransactionById)
	r.GET("/premium", h.GetPremiumList)
	r.PUT("/premium/:id", h.DeserializeUser(), h.UpdateTransactionStatus)
	r.PUT("/transaction/:id", h.DeserializeUser(), h.TransactionUpdate)
	r.POST("/premium/transaction", h.PremiumTransaction)

	r.GET("/user/all-transfer/:id", h.DeserializeUser(), h.GetHistoryTransactionUser)

	r.POST("/send-message", h.PayMessagePost)
	r.GET("/send-message/:id", h.PayMessageGet)

	r.POST("/nft", h.CreateNFT)
	r.GET("/nft/:id", h.DeserializeUser(), h.GetNFTByID)
	r.GET("/nft", h.DeserializeUser(), h.GetNFTList)
	r.PUT("/nft/:id", h.DeserializeUser(), h.UpdateNFT)
	r.DELETE("/nft/:id", h.DeserializeUser(), h.DeleteNFT)

	r.POST("/coin/nft", h.DeserializeUser(), h.CoinNFTCreate)
	r.GET("/coin/nft/:id", h.GetCoinNFTByID)
	r.GET("/coin/nft", h.GetCoinNFTList)
	r.PUT("/coin/nft/:id", h.DeserializeUser(), h.UpdateCoinNFT)
	r.DELETE("/coin/nft/:id", h.DeserializeUser(), h.DeleteCoinNFT)

	r.GET("/success-img/:id", h.GetTransactionSuccessImg)

	url := ginSwagger.URL("swagger/doc.json")
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
