package pattern

import (
	"2-B/internal/contract"
	"fmt"
)

type Single struct {
	Name string
	contract.Pattern
}

func NewSingle(cards []contract.Card) *Single {
	return &Single{
		Pattern: contract.Pattern{
			Cards: cards,
		},
		Name: "單張",
	}
}

func (s *Single) IsGreaterThan(pattern contract.CardPattern) bool {
	single, ok := pattern.(*Single)
	if !ok {
		return false
	}
	return s.Cards[0].GreaterThan(single.Cards[0])
}

func (s *Single) String() string {
	name := s.Cards[0].String()

	return fmt.Sprintf("%s %s", s.Name, name)
}
