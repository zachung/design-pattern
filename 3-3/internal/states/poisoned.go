package states

import (
	"3-3/internal/contract"
)

type Poisoned struct {
	SImpl
	poisoning bool
}

func NewPoisoned(role *contract.Role) *contract.State {
	state := Poisoned{
		SImpl: SImpl{
			name:   "Poisoned",
			rounds: 3,
			role:   role,
		},
	}
	s := contract.State(&state)

	return &s
}

func (s *Poisoned) OnRoundStart(m *contract.Map) {
	s.poisoning = true
	(*s.role).SubHealth(15)
	s.poisoning = false
}

func (s *Poisoned) OnTakeDamage() {
	if s.poisoning {
		return
	}
	(*s.role).SetState(*NewInvincible(s.role))
}
