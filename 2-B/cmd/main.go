package main

import (
	"2-B/internal"
	"2-B/internal/contract"
	"log"
)

var cards = []contract.Card{
	{"S", "A"},
	{"H", "A"},
	{"D", "A"},
	{"C", "A"},
	{"S", "2"},
	{"H", "2"},
	{"D", "2"},
	{"C", "2"},
	{"S", "3"},
	{"H", "3"},
	{"D", "3"},
	{"C", "3"},
	{"S", "4"},
	{"H", "4"},
	{"D", "4"},
	{"C", "4"},
	{"S", "5"},
	{"H", "5"},
	{"D", "5"},
	{"C", "5"},
	{"S", "6"},
	{"H", "6"},
	{"D", "6"},
	{"C", "6"},
	{"S", "7"},
	{"H", "7"},
	{"D", "7"},
	{"C", "7"},
	{"S", "8"},
	{"H", "8"},
	{"D", "8"},
	{"C", "8"},
	{"S", "9"},
	{"H", "9"},
	{"D", "9"},
	{"C", "9"},
	{"S", "10"},
	{"H", "10"},
	{"D", "10"},
	{"C", "10"},
	{"S", "J"},
	{"H", "J"},
	{"D", "J"},
	{"C", "J"},
	{"S", "Q"},
	{"H", "Q"},
	{"D", "Q"},
	{"C", "Q"},
	{"S", "K"},
	{"H", "K"},
	{"D", "K"},
	{"C", "K"},
}
var input = []string{
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
}

func main() {
	log.SetFlags(0)

	big2 := internal.NewBig2()
	big2.SetDeck(internal.NewDeck(cards))
	big2.SetPlayers([4]*internal.Player{
		internal.NewPlayer("水球"),
		internal.NewPlayer("火球"),
		internal.NewPlayer("保齡球"),
		internal.NewPlayer("地瓜球"),
	})
	big2.Start(inputCh())
}

func inputCh() chan string {
	ch := make(chan string)
	go func() {
		for _, s := range input {
			ch <- s
		}
	}()
	return ch
}
