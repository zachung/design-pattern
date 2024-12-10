package states

import (
	"3-3/internal/contract"
)

type Teleport struct {
	SImpl
}

func NewTeleport(role *contract.Role) *contract.State {
	state := Teleport{
		SImpl: SImpl{
			name:   "Teleport",
			rounds: 1,
			role:   role,
		},
	}
	s := contract.State(&state)

	return &s
}

func (s *Teleport) OnRoundEnd(m *contract.Map) {
	s.rounds--
	if s.rounds <= 0 {
		(*m).PutInRandomLocation(contract.Object(*s.role))
		(*s.role).SetState(*NewNormal(s.role))
	}
}
