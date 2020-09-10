package main

import (
	"fmt"
	"sync"
	"time"
)

var WG sync.WaitGroup

func main() {
	// 測試協程同步
	Read()
	go Write()
	// 等待 Write 內的邏輯處理完
	WG.Wait()
	fmt.Println("All done !")
}

// 讀取數據
func Read() {
	for i := 0; i < 3; i++ {
		WG.Add(1)
	}
}

// 寫入數據
func Write() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Done ->", i)
		WG.Done()
	}
}
