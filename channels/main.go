package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// create channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// 每個 go routine 運行一次後就關掉
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// 無限執行, 執行完時再產生新的等待 channel
	// for {
	// 	go checkLink(<-c, c)
	// }

	// 無限執行寫法 2
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// 無限執行, 每個 go routine 子線程加上緩衝時間 5 秒, 以下寫法會中斷主進程
	// for l := range c {
	// 	time.Sleep(5 * time.Second)
	// 	checkLink(l, c)
	// }

	// 無限執行, 每個 go routine 子線程加上緩衝時間 5 秒, 以下寫法不會中斷主進程
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// send message to channel
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// send message to channel
	// c <- "Yep its up"
	c <- link

}
