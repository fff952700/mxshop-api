package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
