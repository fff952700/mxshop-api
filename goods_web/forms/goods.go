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
	ID              int32    `json:"id" form:"id"`
	Name            string   `json:"name" form:"name"`
	GoodsSn         string   `json:"goods_sn" form:"goods_sn"`
	Stocks          int32    `json:"stocks" form:"stocks"`
	MarketPrice     float32  `json:"market_price" form:"market_price"`
	ShopPrice       float32  `json:"shop_price" form:"shop_price"`
	GoodsBrief      string   `json:"goods_brief" form:"goods_brief"`
	ShipFree        bool     `json:"ship_free" form:"ship_free"`
	Images          []string `json:"images" form:"images"`
	DescImages      []string `json:"desc_images"`
	GoodsFrontImage string   `json:"front_image"`
	IsNew           bool     `json:"is_new" form:"is_new"`
	IsHot           bool     `json:"is_hot" form:"is_hot"`
	OnSale          bool     `json:"on_sale" form:"on_sale"`
	CategoryId      int32    `json:"category_id" form:"category_id"`
	BrandId         int32    `json:"brand_id" form:"brand_id"`
}
