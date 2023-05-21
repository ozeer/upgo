package global

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

// 错误打印日志方法
func Error(errMsg string) {
	Logger.Error(errMsg)
	// os.Exit(1)
}
