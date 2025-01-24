package contract

import "fmt"

var ranks = map[string]int{
	"3":  0,
	"4":  1,
	"5":  2,
	"6":  3,
	"7":  4,
	"8":  5,
	"9":  6,
	"10": 7,
	"J":  8,
	"Q":  9,
	"K":  10,
	"A":  11,
	"2":  12,
}
var suits = map[Suit]int{
	"C": 0,
	"D": 1,
	"H": 2,
	"S": 3,
}

type Suit string
type Rank string

type Card struct {
	Suit Suit
	Rank string
}

func (c Card) String() string {
	return fmt.Sprintf("%s[%s]", c.Suit, c.Rank)
}

func (c Card) GreaterThan(another Card) bool {
	if another.Rank != c.Rank {
		return ranks[c.Rank] > ranks[another.Rank]
	}
	return suits[c.Suit] > suits[another.Suit]
}
