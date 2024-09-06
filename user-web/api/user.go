package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"
	"strconv"
	"time"

	"mxshop-api/user-web/proto"
	"net/http"
)

var conn *grpc.ClientConn
var userClient proto.UserClient

func InitUserClient() {
	var err error
	conn, err = grpc.NewClient(fmt.Sprintf("%s:%d", global.ServerConf.UserServerInfo.Host, global.ServerConf.UserServerInfo.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Error("【InitUserClient】连接用户服务失败：", err)
		return
	}
	userClient = proto.NewUserClient(conn)
}

func GetUserList(ctx *gin.Context) {
	// 获取偏移量和分页数量
	pn := ctx.DefaultQuery("pn", "0")
	// 字符串转换uint32
	pnInt, _ := strconv.Atoi(pn)

	pSize := ctx.DefaultQuery("pSize", "25")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		global.HandleGrpcErrToHttp(err, ctx)
		return
	}
	data := mapUserData(rsp)
	ctx.JSON(http.StatusOK, data)
}

func mapUserData(rsp *proto.UserListResponse) []interface{} {
	u := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := make(map[string]interface{}, 0)
		user["id"] = value.Id
		user["mobile"] = value.Mobile
		user["nickname"] = value.Nickname
		user["gender"] = value.Gender
		user["role"] = value.Role
		u = append(u, user)
	}
	return u
}

// 用户登陆
func PassWordLogin(c *gin.Context) {
	// 表单校验
	// 实例化表单
	passWordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&passWordLoginForm); err != nil {
		// 使用翻译器进行翻译
		global.HandlerValidatorError(c, err)
		return
	}
	fmt.Println(passWordLoginForm.Mobile)
	// 通过手机号判断用户是否存在
	rsp, err := userClient.GetUserByMobile(c, &proto.MobileRequest{
		Mobile: passWordLoginForm.Mobile,
	})
	fmt.Printf("user:%v", rsp)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			// 格式化错误成功
			switch e.Code() {
			case codes.InvalidArgument:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "系统错误",
				})
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "用户不存在",
				})
			}
		}
	} else {
		// 校验密码是否正确
		verify, _ := userClient.CheckUserPasswd(c, &proto.PasswordCheckInfo{
			Password:          passWordLoginForm.Password,
			EncryptedPassword: rsp.Password,
		})
		fmt.Println("result:", verify)
		if !verify.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "密码错误",
			})
			return
		}
		// 实例化jwt对象
		j := middlewares.NewJWT()
		// 实例化claims
		claims := models.CustomClaims{
			ID:          uint(rsp.Id),
			NickName:    rsp.Nickname,
			AuthorityId: uint(rsp.Role),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
				Issuer:    "test",
			},
		}
		// 创建token
		token, err := j.CreateToken(claims)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "创建token失败",
			})
		}
		data := map[string]interface{}{
			"Token":     token,
			"ExpiresAt": time.Now().Unix() + 60*60,
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "登陆成功",
			"data": data,
		})
	}
}
