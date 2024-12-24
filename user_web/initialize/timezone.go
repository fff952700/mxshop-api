package initialize

import (
	"time"

	"mxshop_api/user_web/global"
)

// InitTimeZone 设置时区
func init() {
	cst, _ := time.LoadLocation(global.ServerConf.TimeZone)
	global.TimeZone = cst

}
