package test

import (
	"context"
	"os"
	"os/signal"
	"testing"
	"time"
)

// 产生可被中断的 context
func InterruptibleContext() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt)
}

func TestInterruptibleContext(t *testing.T) {
	//在测试中我们使用一个可以手动取消的context
	ctx, cancel := InterruptibleContext()

	go func() {
		time.Sleep(100 * time.Minute) //模拟os.Interrupt的延迟
		cancel()
	}()

	select {
	case <-time.After(101 * time.Minute):
		t.Error("context should have been cancelled")
	case <-ctx.Done():
		t.Log("通过信号结束")
	}
}
