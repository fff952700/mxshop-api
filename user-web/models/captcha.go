package models

type CaptchaBody struct {
	Id          string
	CaptchaType string `json:"captcha_type" binding:"required"`
	VerifyValue string
}
