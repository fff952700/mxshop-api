package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"

	"mxshop_api/user_web/global"
)

func GenerateCaptchaHandler(c *gin.Context) {
	var driver base64Captcha.Driver
	info := global.Cfg.CaptchaInfo

	// Choose the captcha driver based on configuration
	switch info.Type {
	case "audio":
		driver = base64Captcha.NewDriverAudio(6, "en")
	case "string":
		driver = base64Captcha.NewDriverString(80, 240, 4, base64Captcha.OptionShowSineLine, 6, "1234567890abcdefghijklmnopqrstuvwxyz", nil, base64Captcha.DefaultEmbeddedFonts, []string{"wqy-microhei.ttc"})
	case "math":
		driver = base64Captcha.NewDriverMath(80, 240, 4, base64Captcha.OptionShowSineLine, nil, base64Captcha.DefaultEmbeddedFonts, []string{"wqy-microhei.ttc"})
	case "chinese":
		driver = base64Captcha.NewDriverChinese(80, 240, 4, base64Captcha.OptionShowSineLine, 6, info.SourceChinese, nil, base64Captcha.DefaultEmbeddedFonts, []string{"wqy-microhei.ttc"})
	default:
		driver = base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	}

	// Ensure Redis client is initialized
	if global.RedisClient == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": "Redis client not initialized"})
		return
	}

	// Create a new captcha instance
	captcha := base64Captcha.NewCaptcha(driver, global.RedisStore)
	id, _, answer, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	// Return the captcha data
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": answer, "captcha_id": id, "msg": "success"})
}
