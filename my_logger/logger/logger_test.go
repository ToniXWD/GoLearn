package logger_test

import (
	"my_logger/logger" // 假设我们的日志库存在该路径下
	"testing"
)

func Test_MyLogger(t *testing.T) {
	logger.Info("This is a information log")
	logger.Error("This is a error log")
	logger.Debug("This is a debug log")
}
