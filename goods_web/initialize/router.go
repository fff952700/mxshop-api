package initialize

import (
	"github.com/gin-gonic/gin"

	"mxshop_api/goods_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/g/v1")
	router.InitGoodsRouter(ApiGroup)
	return Router
}
