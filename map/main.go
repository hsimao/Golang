package main

import "fmt"

type hello struct {
	name string
	age  int
}

func main() {
	// var colors map[string]string

	// 宣告 map
	// colors := make(map[string]string)

	// 宣告 map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// 增加
	// colors["white"] = "#ffffff"

	// 刪除
	// delete(colors, "white")

	printMap(colors)
}

// loop map
func printMap(c map[string]string) {
	for color, value := range c {
		fmt.Println("Hex code for", color, "is", value)
	}
}
