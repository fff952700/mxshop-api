package router

import (
	"github.com/gin-gonic/gin"

	"mxshop_api/user_web/api"
	"mxshop_api/user_web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userGroup := Router.Group("user")
	{
		userGroup.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		userGroup.POST("/login", api.PassWordLoginForms)
		userGroup.POST("/register", api.RegisterUser)
	}
}
