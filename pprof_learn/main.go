package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func busyWork() {
	for i := 0; i < 1e7; i++ {
		_ = i * i
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// 启动 pprof 服务器
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// 手动启动 CPU 分析
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("could not create CPU profile: ", err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start CPU profile: ", err)
		return
	}
	defer pprof.StopCPUProfile()

	// 模拟一些工作负载
	for i := 0; i < 10; i++ {
		busyWork()
		time.Sleep(1 * time.Second)
	}

	// 手动生成内存分析
	f, err = os.Create("mem.prof")
	if err != nil {
		fmt.Println("could not create memory profile: ", err)
		return
	}
	defer f.Close()
	runtime.GC() // 获取准确的内存分配数据
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("could not write memory profile: ", err)
		return
	}

	wg.Wait()
}
