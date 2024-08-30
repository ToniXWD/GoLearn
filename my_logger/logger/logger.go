package logger

import (
	"fmt"
	"runtime"
	"time"
)

// ANSI color codes
const (
	InfoColor  = "\033[1;32m%s\033[0m"
	ErrorColor = "\033[1;31m%s\033[0m"
	DebugColor = "\033[0;33m%s\033[0m"
)

func Info(msg string) {
	log(InfoColor, "INFO", msg)
}

func Error(msg string) {
	log(ErrorColor, "ERROR", msg)
}

func Debug(msg string) {
	log(DebugColor, "DEBUG", msg)
}

// 通用的日志函数, 提供日志级别和信息作为参数
func log(color string, level string, msg string) {
	var funcName, file string
	var line int

	pc, file, line, ok := runtime.Caller(2) // 向上2层的调用者, 即调用 Info/Error/Debug 的函数

	if ok {
		funcName = runtime.FuncForPC(pc).Name() // 通过PC获取函数名
	} else {
		funcName = "unknow"
		file = "unknow"
		line = 0
	}

	timeStamp := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("%s [%s] %s:%d %s - %s\n", timeStamp, fmt.Sprintf(color, level), file, line, funcName, msg)
}
