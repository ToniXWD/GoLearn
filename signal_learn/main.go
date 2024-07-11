package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 创建一个可以接收信号的 channel
	sigs := make(chan os.Signal, 1)

	signals := []os.Signal{
		syscall.SIGINT, syscall.SIGTERM,
	}

	// notify函数在有信号来的时候，发送信号到channel
	signal.Notify(sigs, signals...)

	// 开启一个goroutine等待信号的到来
	go func() {
		sig := <-sigs
		switch sig {
		case syscall.SIGINT:
			fmt.Println("Receive SIGINT")
		case syscall.SIGTERM:
			fmt.Println("Receive SIGTERM")
		}
		os.Exit(0)
	}()

	fmt.Println("等待信号")
	select {}
}
