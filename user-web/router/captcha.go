package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("/base")
	{
		BaseRouter.GET("/captcha", api.GenerateCaptchaHandler)
	}
}
