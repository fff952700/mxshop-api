package forms

type GoodsFilter struct {
	PriceMin    int32  `json:"price_min,omitempty" form:"price_min,omitempty"`
	PriceMax    int32  `json:"price_max,omitempty" form:"price_max,omitempty"`
	IsHot       bool   `json:"is_hot,omitempty" form:"is_hot,omitempty"`
	IsNew       bool   `json:"is_new,omitempty" form:"is_new,omitempty"`
	IsTab       bool   `json:"is_tab,omitempty" form:"is_tab,omitempty"`
	TopCategory int32  `json:"top_category,omitempty" form:"top_category,omitempty"`
	Pages       int32  `json:"pages,omitempty" form:"pages,omitempty"`
	PagePerNums int32  `json:"page_per_nums,omitempty" form:"page_per_nums,omitempty"`
	KeyWords    string `json:"key_words,omitempty" form:"key_words,omitempty"`
	Brand       int32  `json:"brand,omitempty" form:"brand,omitempty"`
}

type GoodsInfo struct {
	ID              int32    `json:"id" form:"id,omitempty"`
	Name            string   `json:"name" form:"name,omitempty"`
	GoodsSn         string   `json:"goods_sn" form:"goods_sn,omitempty"`
	Stocks          int32    `json:"stocks" form:"stocks,omitempty"`
	MarketPrice     float32  `json:"market_price" form:"market_price,omitempty"`
	ShopPrice       float32  `json:"shop_price" form:"shop_price,omitempty"`
	GoodsBrief      string   `json:"goods_brief" form:"goods_brief,omitempty"`
	ShipFree        bool     `json:"ship_free" form:"ship_free,omitempty"`
	Images          []string `json:"images" form:"images,omitempty"`
	DescImages      []string `json:"desc_images,omitempty"`
	GoodsFrontImage string   `json:"front_image,omitempty"`
	IsNew           bool     `json:"is_new" form:"is_new,omitempty"`
	IsHot           bool     `json:"is_hot" form:"is_hot,omitempty"`
	OnSale          bool     `json:"on_sale" form:"on_sale,omitempty"`
	CategoryId      int32    `json:"category_id" form:"category_id,omitempty"`
	BrandId         int32    `json:"brand_id" form:"brand_id,omitempty"`
}
