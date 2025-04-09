package initialize

import (
	"github.com/gin-gonic/gin"

	"mxshop_api/order_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	ApiGroup := Router.Group("/o/v1")

	router.InitShopCartRouter(ApiGroup)
	return Router
}
