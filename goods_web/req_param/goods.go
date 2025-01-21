package req_param

type GoodsFilter struct {
	PMin     int    `forms:"pmin" json:"p_min,omitempty"`
	PMax     int    `forms:"pmax" json:"p_max,omitempty"`
	IsHot    int    `forms:"ih" json:"is_hot,omitempty"`
	IsNew    int    `forms:"in" json:"is_new,omitempty"`
	IsTab    int    `forms:"it" json:"is_tab,omitempty"`
	Category int    `forms:"c" json:"category,omitempty"`
	PN       int    `forms:"pn" json:"pn,omitempty"`
	PNum     int    `forms:"pnum" json:"p_num,omitempty"`
	Keyword  string `forms:"q" json:"keyword,omitempty"`
	Brand    int    `forms:"b" json:"brand,omitempty"`
}
