package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"
	"strconv"
	"time"

	"mxshop-api/user-web/proto"
	"net/http"
)

func GetUserList(ctx *gin.Context) {
	// 获取偏移量和分页数量
	pn := ctx.DefaultQuery("pn", "0")
	// 字符串转换uint32
	pnInt, _ := strconv.Atoi(pn)

	pSize := ctx.DefaultQuery("pSize", "25")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := global.UserClient.GetUserList(context.Background(), &proto.PageInfo{
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

func PassWordLogin(c *gin.Context) {
	passWordLoginForm := forms.PassWordLoginForm{}
	rsp, err := CheckUserForms(c, &passWordLoginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// 校验密码是否正确
	verify, err := global.UserClient.CheckUserPasswd(c, &proto.PasswordCheckInfo{
		Password:          passWordLoginForm.Password,
		EncryptedPassword: rsp.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "密码校验失败"})
		return
	}
	if !verify.Success {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "密码错误"})
		return
	}
	data, err := UserCreateToken(c, rsp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "data": data})
}

func CheckUserForms(c *gin.Context, passWordLoginForm *forms.PassWordLoginForm) (*proto.UserInfoResponse, error) {
	// 表单校验
	if err := c.ShouldBind(passWordLoginForm); err != nil {
		global.HandlerValidatorError(c, err)
		return nil, fmt.Errorf("表单校验失败")
	}
	// 验证码校验
	if global.ServerConf.CaptChaInfo.EnableCaptcha {
		if !global.RedisStore.Verify(passWordLoginForm.CaptchaId, passWordLoginForm.Captcha, true) {
			return nil, fmt.Errorf("验证码错误")
		}
	}

	// 通过手机号判断用户是否存在
	rsp, err := global.UserClient.GetUserByMobile(c, &proto.MobileRequest{
		Mobile: passWordLoginForm.Mobile,
	})
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}

	return rsp, nil
}

func UserCreateToken(c *gin.Context, rsp *proto.UserInfoResponse) (map[string]interface{}, error) {
	// 实例化jwt对象
	j := middlewares.NewJWT()
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
		return nil, fmt.Errorf("创建token失败")
	}

	// 返回token信息
	data := map[string]interface{}{
		"Token":     token,
		"ExpiresAt": time.Now().Unix() + 60*60,
	}
	return data, nil
}
