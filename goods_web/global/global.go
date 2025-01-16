package global

import (
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"mxshop_api/goods_web/config"
	"mxshop_api/goods_web/proto"
)

var (
	Cfg          = &config.Cfg{}
	NacosConf    = &config.NacosConfig{}
	Translator   ut.Translator
	GoodsClient  proto.GoodsClient
	TimeZone     *time.Location
	ConsulClient *api.Client
)

// HandleGrpcErrToHttp 将grpc状态码转换为http
func HandleGrpcErrToHttp(err error, ctx *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
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

func MapToJSONString(fields map[string]string) (string, error) {
	jsonData, err := json.Marshal(fields)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
