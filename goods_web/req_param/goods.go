package req_param

type GoodsFilter struct {
	PMin     int    `form:"pmin" json:"p_min,omitempty"`
	PMax     int    `form:"pmax" json:"p_max,omitempty"`
	IsHot    int    `form:"ih" json:"is_hot,omitempty"`
	IsNew    int    `form:"in" json:"is_new,omitempty"`
	IsTab    int    `form:"it" json:"is_tab,omitempty"`
	Category int    `form:"c" json:"category,omitempty"`
	PN       int    `form:"pn" json:"pn,omitempty"`
	PNum     int    `form:"pnum" json:"p_num,omitempty"`
	Keyword  string `form:"q" json:"keyword,omitempty"`
	Brand    int    `form:"b" json:"brand,omitempty"`
}
