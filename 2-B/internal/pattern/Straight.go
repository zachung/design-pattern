package pattern

import (
	"2-B/internal/contract"
	"fmt"
)

type Straight struct {
	Name string
	contract.Pattern
}

func NewStraight(cards []contract.Card) *Straight {
	return &Straight{
		Pattern: contract.Pattern{
			Cards: cards,
		},
		Name: "順子",
	}
}

func (s *Straight) IsGreaterThan(pattern contract.CardPattern) bool {
	straight, ok := pattern.(*Straight)
	if !ok {
		return false
	}
	return s.Cards[4].GreaterThan(straight.Cards[4])
}

func (s *Straight) String() string {
	name := s.Cards[0].String()
	for i := 1; i < len(s.Cards); i++ {
		name += " " + s.Cards[i].String()
	}

	return fmt.Sprintf("%s %s", s.Name, name)
}
