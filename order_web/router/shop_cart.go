package router

import (
	"github.com/gin-gonic/gin"

	"mxshop_api/order_web/api/shop_cart"
	"mxshop_api/order_web/middlewares"
)

func InitShopCartRouter(Router *gin.RouterGroup) {
	ShopCartRouter := Router.Group("shopCart")
	{
		ShopCartRouter.GET("getShopCartList", middlewares.JWTAuth(), shop_cart.List)
		ShopCartRouter.POST("createShopCartItem", middlewares.JWTAuth(), shop_cart.New)
		ShopCartRouter.PUT("updateShopCartItem", middlewares.JWTAuth(), shop_cart.Update)
		ShopCartRouter.DELETE("deleteShopCartItem", middlewares.JWTAuth(), shop_cart.Delete)
	}
}
