package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades ♠", "Diamonds ♦", "Hearts ♥", "Clubs ♣"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// ["Spades", "Diamonds"] => "Spades,Diamonds"
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// 寫入檔案
func (d deck) saveToFile(filename string) error {
	// 0666 任何人都可讀寫此檔案
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// 讀取檔案
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// []byte => string => []string
	// [12 23 34 123 34 11] => "Ace of Spades,Two of Spades" => ["Ace of Spades", "Two of Spade"]
	s := strings.Split(string(bs), ",")

	return deck(s)
}

// 洗牌
func (d deck) shuffle() {
	// 必須要更換 rand 的 source 才能每次執行都產生不一樣的亂數
	// 使用 time.Now() 取得當下時間, 並用 UnixNano 取得當下時間的 int64 值
	// 並拿來當 rand.NewSource 的 seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// 陣列位置交換 [1, 3] => [3, 1]
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
