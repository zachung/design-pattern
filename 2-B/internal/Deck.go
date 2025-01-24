package internal

type Deck struct {
	cards []Card
}

func NewDeck(cards []Card) *Deck {
	return &Deck{cards}
}

func (d *Deck) Deal() (card Card) {
	l := len(d.cards) - 1
	card, d.cards = d.cards[l], d.cards[:l]
	return
}

func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}
