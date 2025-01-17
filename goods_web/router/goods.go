package router

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/goods_web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsGroup := Router.Group("goods")
	{
		GoodsGroup.GET("/list", goods.List)
		GoodsGroup.GET("/detail", goods.Detail)
	}
}
