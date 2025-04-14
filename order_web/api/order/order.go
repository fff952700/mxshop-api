package order

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"

	"mxshop_api/order_web/api"
	"mxshop_api/order_web/forms"
	"mxshop_api/order_web/global"
	"mxshop_api/order_web/models"
	"mxshop_api/order_web/proto"
)

func List(ctx *gin.Context) {
	// 获取用户id
	var (
		orderScopesForm forms.OrderScopesForm
	)
	userId := api.GetUserId(ctx)
	if err := ctx.ShouldBindBodyWithJSON(&orderScopesForm); err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	// 创建统一基础请求
	req := &proto.OrderFilterRequest{
		Pages:       orderScopesForm.Page,
		PagePerNums: orderScopesForm.PageNum,
	}
	// 判断是否为管理员
	claims, _ := ctx.Get("claims")
	model := claims.(*models.CustomClaims)
	if model.AuthorityId != 1 {
		req.UserId = userId
	}
	rsp, err := global.OrderClient.OrderList(ctx, req)
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total": rsp.Total, "data": rsp.Data})
}

func New(ctx *gin.Context) {
	// 获取传入的购物车参数
	var orderRequest forms.OrderRequestForm
	if err := ctx.ShouldBind(&orderRequest); err != nil {
		zap.S().Errorw("参数错误",
			"msg", err.Error(),
		)
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	userId := api.GetUserId(ctx)
	if userId == 0 {
		zap.S().Infof("用户不存在")
		return
	}
	resp, err := global.OrderClient.CreateOrder(ctx, &proto.OrderRequest{
		UserId:  userId,
		Name:    orderRequest.Name,
		Mobile:  orderRequest.Mobile,
		Address: orderRequest.Address,
		Post:    orderRequest.Post,
	})
	if err != nil {
		zap.S().Error(err)
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})

}

func Detail(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {}
