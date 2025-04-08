package initialize

import (
	"github.com/gin-gonic/gin"

	"mxshop_api/order_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/g/v1")
	router.InitShopCartRouter(ApiGroup)
	return Router
}
