package states

import (
	"3-3/internal/contract"
)

type Erupting struct {
	SImpl
}

func NewErupting(role *contract.Role) *contract.State {
	state := Erupting{
		SImpl: SImpl{
			name:   "Erupting",
			rounds: 3,
			role:   role,
		},
	}
	s := contract.State(&state)

	return &s
}

func (s *Erupting) OnRoundEnd(*contract.Map) {
	s.rounds--
	if s.rounds <= 0 {
		(*s.role).SetState(*NewTeleport(s.role))
	}
}

func (s *Erupting) OnAttack(m *contract.Map) {
	// attack all roles
	roles := (*m).GetRoles()
	for _, role := range roles {
		if *role == *s.role {
			// ignore self
			continue
		}
		(*role).SubHealth(50)
	}
}
