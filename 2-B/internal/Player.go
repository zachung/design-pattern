package internal

import "slices"

type Player struct {
	Name  string
	Cards []Card
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Cards: make([]Card, 0),
	}
}

func (p *Player) AddCard(card Card) {
	var i int
	for i = 0; i < len(p.Cards); i++ {
		if !p.Cards[i].GreaterThan(card) {
			break
		}
	}
	p.Cards = slices.Insert(p.Cards, i, card)
}

func (p *Player) IsHandEmpty() bool {
	return len(p.Cards) == 0
}

func (p *Player) Play(cardIndexes []int) *CardPattern {
	if cardIndexes[0] == -1 {
		return nil
	}
	// 挑出手牌
	var cards []Card
	var less []Card
out:
	for i, card := range p.Cards {
		for _, cardIndex := range cardIndexes {
			if cardIndex == i {
				cards = append(cards, p.Cards[i])
				continue out
			}
		}
		less = append(less, card)
	}
	p.Cards = less

	return NewCardPattern(cards)
}
