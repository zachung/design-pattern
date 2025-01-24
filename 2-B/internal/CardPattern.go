package internal

import (
	"fmt"
)

type CardPattern struct {
	cards []Card
}

func NewCardPattern(cards []Card) *CardPattern {
	return &CardPattern{cards: cards}
}

func (p CardPattern) String() string {
	name := p.cards[0].String()
	for i := 1; i < len(p.cards); i++ {
		name += " " + p.cards[i].String()
	}
	// TODO: 檢測牌型
	patternName := "單張"

	return fmt.Sprintf("%s %s", patternName, name)
}
