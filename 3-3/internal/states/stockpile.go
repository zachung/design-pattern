package states

import (
	"3-3/internal/contract"
	"fmt"
)

type Stockpile struct {
	SImpl
}

func NewStockpile(role *contract.Role) *contract.State {
	state := Stockpile{
		SImpl: SImpl{
			name:   "Stockpile",
			rounds: 2,
			role:   role,
		},
	}
	fmt.Printf("new %p %T\n", role, role)
	s := contract.State(&state)

	return &s
}

func (s *Stockpile) OnTakeDamage() {
	(*s.role).SetState(*NewNormal(s.role))
}

func (s *Stockpile) OnRoundEnd(*contract.Map) {
	s.rounds--
	if s.rounds <= 0 {
		(*s.role).SetState(*NewErupting(s.role))
	}
}
