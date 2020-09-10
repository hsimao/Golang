package main

import (
	"fmt"
	"time"
)

var chanInt chan int = make(chan int, 10)
var timeout chan bool = make(chan bool)

func main() {
	// 啟動發送數據的協程
	go Send()
	// 啟動接收數據的協程
	go Receive()
	time.Sleep(time.Second * 30)
}

func Loop() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Microsecond * 10)
		fmt.Printf("%d,", i)
	}
}

// 發送數據到 chan
func Send() {
	time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
	time.Sleep(time.Second * 2)
	timeout <- true
}

// 接收 chan 中的數據
func Receive() {
	// num := <-chanInt
	// fmt.Println("mun:", num)
	// num = <-chanInt
	// fmt.Println("mun:", num)
	// num = <-chanInt
	// fmt.Println("mun:", num)

	for {
		// 不停的從 chan 中讀取數據
		select {
		case num := <-chanInt:
			fmt.Println("mun", num)
		case <-timeout:
			fmt.Println("timeout...")
		}
	}
}
