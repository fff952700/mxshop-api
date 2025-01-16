package initialize

import (
	"github.com/gin-gonic/gin"

	"mxshop_api/goods_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	ApiGroup := Router.Group("/g/v1")
	router.InitGoodsRouter(ApiGroup)
	return Router
}
