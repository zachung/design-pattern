package states

import (
	"3-3/internal/contract"
	"fmt"
)

type SImpl struct {
	role   *contract.Role
	name   string
	rounds int
}

func (s *SImpl) GetStateName() string {
	return fmt.Sprintf("%s(%d)", s.name, s.rounds)
}

func (s *SImpl) OnRoundStart(*contract.Map) {
}

func (s *SImpl) OnRounding(*contract.Map) {
	(*s.role).DoAction()
}

func (s *SImpl) OnRoundEnd(*contract.Map) {
	s.rounds--
	if s.rounds <= 0 {
		(*s.role).SetState(*NewNormal(s.role))
	}
}

func (s *SImpl) OnTakeDamage() {
	(*s.role).SetState(*NewInvincible(s.role))
}

func (s *SImpl) OnAttack(*contract.Map) {
	(*s.role).Attack()
}

func (s *SImpl) GetAllowDirections() map[rune]rune {
	return contract.Directions
}
