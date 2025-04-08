package shop_cart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mxshop_api/order_web/api"
	"mxshop_api/order_web/forms"
	"mxshop_api/order_web/global"
	"mxshop_api/order_web/proto"
)

func List(ctx *gin.Context) {
	userId := api.GetUserId(ctx)
	if userId == 0 {
		return
	}
	// 调用order服务
	resp, err := global.OrderClient.CartItemList(ctx, &proto.UserInfo{UserId: userId})
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total": resp.Total, "data": resp.Data})

}

func New(ctx *gin.Context) {
	// 获取传入的购物车参数
	var cartItem forms.ShopCartItemForm
	if err := ctx.ShouldBindJSON(&cartItem); err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	userId := api.GetUserId(ctx)
	if userId == 0 {
		return
	}
	// 调用order服务
	resp, err := global.OrderClient.CreateCartItem(ctx, &proto.CartItemRequest{UserId: userId, GoodsId: cartItem.GoodsId, Nums: cartItem.Nums, Checked: cartItem.Checked})
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": resp.Id, "msg": "success"})

}

func Update(ctx *gin.Context) {
	// 获取传入的购物车参数
	var cartItem forms.ShopCartItemUpdateForm
	if err := ctx.ShouldBindJSON(&cartItem); err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	userId := api.GetUserId(ctx)
	if userId == 0 {
		return
	}
	_, err := global.OrderClient.UpdateCartItem(ctx, &proto.CartItemRequest{Id: cartItem.Id, UserId: userId, Nums: cartItem.Nums, Checked: cartItem.Checked})
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})

}

func Delete(ctx *gin.Context) {
	// 获取传入的购物车参数
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	_, err = global.OrderClient.DeleteCartItem(ctx, &proto.CartItemRequest{Id: int32(i)})
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}
