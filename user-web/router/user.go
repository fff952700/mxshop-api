package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userGroup := Router.Group("user")
	{
		userGroup.GET("/list", api.GetUserList)
	}
}
