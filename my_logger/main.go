package main

import (
	"fmt"
	"runtime"
)

func main() {
	PrintCallerInfo()
}

func PrintCallerInfo() {
	// skip设置为1，将跳过PrintCallerInfo的帧，获取PrintCallerInfo调用者（此处为main）的帧信息
	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	funcName := runtime.FuncForPC(pc).Name()

	fmt.Printf("调用函数: %s, 文件: %s, 行号: %d, PC值: %v\n", funcName, file, line, pc)
}
