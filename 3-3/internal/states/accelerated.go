package states

import (
	"3-3/internal/contract"
)

type Accelerated struct {
	SImpl
}

func NewAccelerated(role *contract.Role) *contract.State {
	state := Accelerated{
		SImpl: SImpl{
			name:   "Accelerated",
			rounds: 3,
			role:   role,
		},
	}
	s := contract.State(&state)

	return &s
}

func (s *Accelerated) OnRounding(m *contract.Map) {
	(*s.role).DoAction()
	(*m).Draw()
	(*s.role).DoAction()
}

func (s *Accelerated) OnTakeDamage() {
	(*s.role).SetState(*NewNormal(s.role))
}
