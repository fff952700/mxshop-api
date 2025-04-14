package router

import (
	"github.com/gin-gonic/gin"

	"mxshop_api/order_web/api/order"
	"mxshop_api/order_web/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders")
	{
		OrderRouter.GET("getList", middlewares.JWTAuth(), order.List)
		OrderRouter.POST("add", middlewares.JWTAuth(), order.New)
	}
}
