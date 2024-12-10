package states

import "3-3/internal/contract"

type Invincible struct {
	SImpl
}

func NewInvincible(role *contract.Role) *contract.State {
	state := Invincible{
		SImpl: SImpl{
			name:   "Invincible",
			rounds: 2,
			role:   role,
		},
	}
	s := contract.State(&state)

	return &s
}
