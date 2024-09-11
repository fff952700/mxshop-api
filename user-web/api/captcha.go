package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"mxshop-api/user-web/global"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GenerateCaptchaHandler(c *gin.Context) {
	var driver base64Captcha.Driver
	// 获取captcha type
	info := global.ServerConf.CaptChaInfo
	// Switch based on captcha type in configuration
	switch info.Type {
	case "audio":
		// Initialize DriverAudio
		driverAudio := base64Captcha.NewDriverAudio(6, "en")
		driver = driverAudio
	case "string":
		// Initialize DriverString if needed
		driverString := &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      4,
			ShowLineOptions: base64Captcha.OptionShowSineLine,
			Length:          6,
			Source:          "1234567890abcdefghijklmnopqrstuvwxyz",
			Fonts:           []string{"wqy-microhei.ttc"}, // Ensure the font file exists
		}
		driver = driverString.ConvertFonts()
	case "math":
		// Initialize DriverMath
		driver = base64Captcha.NewDriverMath(
			80,
			240,
			4,
			base64Captcha.OptionShowSineLine,
			nil,
			base64Captcha.DefaultEmbeddedFonts,
			[]string{"wqy-microhei.ttc"},
		)
	case "chinese":
		// Initialize DriverChinese
		driverChinese := base64Captcha.NewDriverChinese(
			80,                               // height
			240,                              // width
			4,                                // noise count
			base64Captcha.OptionShowSineLine, // show line options
			6,                                // captcha length
			global.ServerConf.CaptChaInfo.SourceChinese, // source characters
			nil,                                // background color (nil for default)
			base64Captcha.DefaultEmbeddedFonts, // fonts storage
			[]string{"wqy-microhei.ttc"},       // fonts
		)
		driver = driverChinese
	default:
		// Default to DriverDigit
		driverDigit := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
		driver = driverDigit
	}
	// generate captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	// respond with captcha id and base64 image
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": b64s, "captcha_id": id, "captcha": answer, "msg": "success"})
}
