package models

import "github.com/mojocn/base64Captcha"

type CaptchaBody struct {
	Id            string
	CaptchaType   string `json:"captcha_type" binding:"required"`
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}
