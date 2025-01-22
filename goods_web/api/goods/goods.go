package goods

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"mxshop_api/goods_web/forms"
	"mxshop_api/goods_web/global"
	"mxshop_api/goods_web/proto"
)

func List(ctx *gin.Context) {
	var goodsForm forms.GoodsFilter
	if err := ctx.ShouldBindQuery(&goodsForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}
	request := &proto.GoodsFilterRequest{
		PriceMin:    goodsForm.PriceMin,
		PriceMax:    goodsForm.PriceMax,
		IsHot:       goodsForm.IsHot,
		IsNew:       goodsForm.IsNew,
		IsTab:       goodsForm.IsTab,
		TopCategory: goodsForm.TopCategory,
		Pages:       goodsForm.Pages,
		PagePerNums: goodsForm.PagePerNums,
		KeyWords:    goodsForm.KeyWords,
		Brand:       goodsForm.Brand,
	}
	zap.S().Infof("request: %v", request)

	response, err := global.GoodsClient.GoodsList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("Error", "err", err.Error())
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	responseMap := map[string]interface{}{
		"total": response.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, goods := range response.Data {
		goodsMap := map[string]interface{}{
			"id":          goods.Id,
			"category_id": goods.CategoryId,
			"name":        goods.Name,
			"goods_brief": goods.GoodsBrief,
			"desc":        goods.GoodsDesc,
			"ship_free":   goods.ShipFree,
			"desc_image":  goods.DescImages,
			"front_image": goods.GoodsFrontImage,
			"shop_price":  goods.ShopPrice,
			"is_host":     goods.IsHot,
			"is_new":      goods.IsNew,
			"on_sale":     goods.OnSale,
		}
		goodsList = append(goodsList, goodsMap)
	}
	responseMap["data"] = goodsList
	ctx.JSON(http.StatusOK, responseMap)
}

func Detail(ctx *gin.Context) {
	// 获取商品id
	id := ctx.Query("id")
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	r, err := global.GoodsClient.GetGoodsDetail(ctx, &proto.GoodInfoRequest{Id: int32(i)})
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	rsp := map[string]interface{}{
		"id":          r.Id,
		"category_id": r.CategoryId,
		"name":        r.Name,
		"goods_brief": r.GoodsBrief,
		"desc":        r.GoodsDesc,
		"ship_free":   r.ShipFree,
		"desc_image":  r.DescImages,
		"front_image": r.GoodsFrontImage,
		"shop_price":  r.ShopPrice,
		"is_host":     r.IsHot,
		"is_new":      r.IsNew,
		"on_sale":     r.OnSale,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func Delete(ctx *gin.Context) {
	id := ctx.Query("id")
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}
	r, err := global.GoodsClient.GetGoodsDetail(ctx, &proto.GoodInfoRequest{Id: int32(i)})
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": r.Id, "msg": "success"})
}

func New(ctx *gin.Context) {
	var goodsInfo forms.GoodsInfo
	if err := ctx.ShouldBindJSON(&goodsInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}
	req := &proto.CreateGoodsInfo{
		CategoryId:      goodsInfo.CategoryId,
		BrandId:         goodsInfo.BrandId,
		OnSale:          goodsInfo.OnSale,
		ShipFree:        goodsInfo.ShipFree,
		IsNew:           goodsInfo.IsNew,
		IsHot:           goodsInfo.IsHot,
		Name:            goodsInfo.Name,
		GoodsSn:         goodsInfo.GoodsSn,
		MarketPrice:     goodsInfo.MarketPrice,
		ShopPrice:       goodsInfo.ShopPrice,
		GoodsBrief:      goodsInfo.GoodsBrief,
		GoodsFrontImage: goodsInfo.GoodsFrontImage,
		Images:          goodsInfo.Images,
		DescImages:      goodsInfo.DescImages,
	}
	rsp, err := global.GoodsClient.CreateGoods(context.Background(), req)
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": rsp.Id, "msg": "success"})
}

func Update(ctx *gin.Context) {
	var goodsInfo forms.GoodsInfo
	if err := ctx.ShouldBindJSON(&goodsInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}
	req := &proto.CreateGoodsInfo{
		Id:              goodsInfo.ID,
		CategoryId:      goodsInfo.CategoryId,
		BrandId:         goodsInfo.BrandId,
		OnSale:          goodsInfo.OnSale,
		ShipFree:        goodsInfo.ShipFree,
		IsNew:           goodsInfo.IsNew,
		IsHot:           goodsInfo.IsHot,
		Name:            goodsInfo.Name,
		GoodsSn:         goodsInfo.GoodsSn,
		MarketPrice:     goodsInfo.MarketPrice,
		ShopPrice:       goodsInfo.ShopPrice,
		GoodsBrief:      goodsInfo.GoodsBrief,
		GoodsFrontImage: goodsInfo.GoodsFrontImage,
		Images:          goodsInfo.Images,
		DescImages:      goodsInfo.DescImages,
	}
	if _, err := global.GoodsClient.UpdateGoods(context.Background(), req); err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": req.Id, "msg": "success"})
}
