package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"mxshop_api/order_web/global"
)

func GetUserId(c *gin.Context) int32 {
	userId, exists := c.Get("userId")
	if !exists {
		global.HandleGrpcErrToHttp(status.Errorf(codes.InvalidArgument, "用户不存在"), c)
		return 0
	}
	return userId.(int32)

}
