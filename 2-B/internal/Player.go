package internal

import (
	"2-B/internal/contract"
	"slices"
)

type Player struct {
	Name  string
	Cards []contract.Card
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Cards: make([]contract.Card, 0),
	}
}

func (p *Player) AddCard(card contract.Card) {
	var i int
	for i = 0; i < len(p.Cards); i++ {
		if p.Cards[i].GreaterThan(card) {
			break
		}
	}
	p.Cards = slices.Insert(p.Cards, i, card)
}

func (p *Player) RemoveCard(card contract.Card) {
	for i, c := range p.Cards {
		if c.String() == card.String() {
			p.Cards = append(p.Cards[:i], p.Cards[i+1:]...)
			return
		}
	}
}

func (p *Player) IsHandEmpty() bool {
	return len(p.Cards) == 0
}

func (p *Player) Play(cardIndexes []int) []contract.Card {
	slices.Sort(cardIndexes)
	// 挑出手牌
	var cards []contract.Card
	for _, index := range cardIndexes {
		cards = append(cards, p.Cards[index])
	}

	return cards
}
