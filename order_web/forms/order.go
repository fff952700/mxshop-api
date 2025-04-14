package forms

type OrderScopesForm struct {
	Page    int32 `json:"page" binding:"required"`
	PageNum int32 `json:"page_num" binding:"required"`
}

type OrderRequestForm struct {
	Address string `json:"address" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Mobile  string `json:"mobile" binding:"required"`
	Post    string `json:"post"`
}
