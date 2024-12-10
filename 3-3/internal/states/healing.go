package states

import (
	"3-3/internal/contract"
)

type Healing struct {
	SImpl
}

func NewHealing(role *contract.Role) *contract.State {
	state := Healing{
		SImpl: SImpl{
			name:   "Healing",
			rounds: 5,
			role:   role,
		},
	}
	s := contract.State(&state)

	return &s
}

func (s *Healing) OnRoundStart(m *contract.Map) {
	r := *s.role
	r.AddHealth(15)
	if r.GetHealth() == r.GetMaxHealth() {
		r.SetState(*NewNormal(s.role))
	}
}
