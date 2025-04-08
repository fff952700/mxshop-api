package forms

type ShopCartItemForm struct {
	GoodsId int32 `json:"goods" binding:"required"`
	Nums    int32 `json:"nums" binding:"required,min=1"`
	Checked bool  `json:"checked" binding:"required"`
}

type ShopCartItemUpdateForm struct {
	Id      int32 `json:"id" binding:"required"`
	Nums    int32 `json:"nums" binding:"required,min=1"`
	Checked bool  `json:"checked" binding:"required"`
}
