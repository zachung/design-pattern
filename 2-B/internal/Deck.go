package internal

import "2-B/internal/contract"

type Deck struct {
	cards []contract.Card
}

func NewDeck(cards []contract.Card) *Deck {
	return &Deck{cards}
}

func (d *Deck) Deal() (card contract.Card) {
	l := len(d.cards) - 1
	card, d.cards = d.cards[l], d.cards[:l]
	return
}

func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}
