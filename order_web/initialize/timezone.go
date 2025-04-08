package initialize

import (
	"time"

	"mxshop_api/order_web/global"
)

// InitTimeZone 设置时区
func init() {
	cst, _ := time.LoadLocation(global.Cfg.ServerInfo.TimeZone)
	global.TimeZone = cst

}
