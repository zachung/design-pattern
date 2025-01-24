package internal

import (
	"2-B/internal/contract"
	"2-B/internal/pattern"
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

func (p *Player) IsHandEmpty() bool {
	return len(p.Cards) == 0
}

func (p *Player) Play(cardIndexes []int) contract.CardPattern {
	// 挑出手牌
	var cards []contract.Card
	var less []contract.Card
	preCards := p.Cards
out:
	for i, card := range preCards {
		for _, cardIndex := range cardIndexes {
			if cardIndex == i {
				cards = append(cards, p.Cards[i])
				continue out
			}
		}
		less = append(less, card)
	}
	p.Cards = less

	switch len(cards) {
	case 1:
		return pattern.NewSingle(cards)
	case 2:
		return pattern.NewPair(cards)
	case 5:
		return pattern.NewStraight(cards)
	}
	p.Cards = preCards

	return nil
}
