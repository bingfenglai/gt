package initialization

import (
	"github.com/bingfenglai/gt/handler"
	"github.com/bingfenglai/gt/router"
)

func InitLogConfig() {

	adaptGinLogToZap()
}

// 将gin日志使用zap输出
func adaptGinLogToZap() {
	router.R.Use(handler.GinZapLogger(), handler.GinZapRecovery(true))

}
