package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("cpu num = %d", runtime.NumCPU())
	// 依據當下系統 cpu 總數, 扣除一個, 保留給維持服務運行用, 避免其他邏輯吃光 cpu 支援
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	// 併發, 以下兩個 func 同時執行
	go Loop()
	go Loop()
	time.Sleep(time.Second * 60)
}

func Loop() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Microsecond * 10)
		fmt.Printf("%d,", i)
	}
}
