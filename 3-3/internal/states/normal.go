package states

import (
	"3-3/internal/contract"
	"fmt"
)

type Normal struct {
	SImpl
}

func NewNormal(role *contract.Role) *contract.State {
	state := Normal{
		SImpl: SImpl{
			name:   "Normal",
			rounds: 1,
			role:   role,
		},
	}
	s := contract.State(&state)

	return &s
}

func (s *Normal) GetStateName() string {
	return fmt.Sprintf("%s", s.name)
}

func (s *Normal) OnRoundEnd(*contract.Map) {
}
