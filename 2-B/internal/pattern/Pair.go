package pattern

import (
	"2-B/internal/contract"
	"fmt"
)

type Pair struct {
	Name string
	contract.Pattern
}

func NewPair(cards []contract.Card) *Pair {
	return &Pair{
		Pattern: contract.Pattern{
			Cards: cards,
		},
		Name: "對子",
	}
}

func (s *Pair) IsGreaterThan(pattern contract.CardPattern) bool {
	pair, ok := pattern.(*Pair)
	if !ok {
		return false
	}
	return s.Cards[1].GreaterThan(pair.Cards[1])
}

func (s *Pair) String() string {
	name := s.Cards[0].String()
	for i := 1; i < len(s.Cards); i++ {
		name += " " + s.Cards[i].String()
	}

	return fmt.Sprintf("%s %s", s.Name, name)
}
