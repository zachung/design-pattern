package states

import (
	"3-3/internal/contract"
	"math/rand/v2"
)

type Orderless struct {
	SImpl
}

func NewOrderless(role *contract.Role) *contract.State {
	state := Orderless{
		SImpl: SImpl{
			name:   "Orderless",
			rounds: 3,
			role:   role,
		},
	}
	s := contract.State(&state)

	return &s
}

func (s *Orderless) OnRounding(*contract.Map) {
	(*s.role).Move()
}

func (s *Orderless) GetAllowDirections() map[rune]rune {
	if rand.IntN(2) == 0 {
		// limit direction
		return map[rune]rune{
			'a': contract.CharacterLeftSymbol,
			'd': contract.CharacterRightSymbol,
		}
	} else {
		return map[rune]rune{
			'w': contract.CharacterUpSymbol,
			's': contract.CharacterDownSymbol,
		}
	}
}
