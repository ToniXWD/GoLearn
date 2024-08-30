package contextlearn_test

import (
	"context"
	"sync"
	"testing"
	"time"
)

// context.WithCancel 返回一个新的 Context 和一个函数，此函数当被调用时，会取消该 Context，释放所有与它关联的资源。
func TestContextWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	select {
	case <-time.After(10 * time.Second):
		t.Error("function did not cancel in expected time")
	case <-ctx.Done():
		// Pass the test
	}
}

// Done channel 将在当前时间加上传入的时间后关闭
func TestContextWithDeadline(t *testing.T) {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		t.Error("function exceeded deadline")
	case <-ctx.Done():
		// Pass the test
	}
}

// context.WithTimeout 是 context.WithDeadline(parent, time.Now().Add(timeout)) 的简化写法。
func TestContextWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		t.Error("function exceeded timeout")
	case <-ctx.Done():
		// Pass the test
	}
}

type myKey struct{}

// context.WithValue 返回一个新的 Context，这个新的 Context 承载了键值对信息，可以通过 context.Value() 方法获取。
func TestContextWithValue(t *testing.T) {
	valCtx := context.WithValue(context.Background(), myKey{}, "my-value")

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		value := valCtx.Value(myKey{})

		if value != "my-value" {
			t.Errorf("got %v, want %v", value, "my-value")
		} else {
			t.Logf("Receive: %s", value)
		}
	}(valCtx)
	wg.Wait()
}
