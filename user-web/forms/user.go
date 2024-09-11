package forms

type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=6,max=6"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required,min=20,max=20"`
}
