package global

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/proto"
	"net/http"
	"strings"
)

var (
	ServerConf *config.ServerConfig = &config.ServerConfig{}
	Translator ut.Translator
	// RedisClient redis客户端
	RedisClient *redis.Client
	// RedisStore redis存储桶
	RedisStore base64Captcha.Store
	UserClient proto.UserClient
)

type JWTInfo struct {
	SigningKey string
}

func HandleGrpcErrToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{"msg": e.Message()})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "用户服务不可用"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "其他错误"})
			}
			return
		}
	}
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func HandlerValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 翻译错误
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"error": RemoveTopStruct(errs.Translate(Translator)),
	})
}
