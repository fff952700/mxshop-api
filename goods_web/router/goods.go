package router

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/goods_web/api/goods"
	"mxshop_api/goods_web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsGroup := Router.Group("goods")
	{
		GoodsGroup.GET("/list", goods.List)
		GoodsGroup.GET("/:id", goods.Detail)
		GoodsGroup.DELETE("/:id", goods.Delete)
		GoodsGroup.POST("/create", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)
	}
}
