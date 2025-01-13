package state

import "3B/internal/contract"

type Poisoned struct {
	name  string
	round int
	role  contract.Role
}

func NewPoisoned(role contract.Role) *Poisoned {
	return &Poisoned{name: "中毒", round: 3, role: role}
}

func (s *Poisoned) GetName() string {
	return s.name
}

func (s *Poisoned) CanAction() bool {
	return true
}

func (s *Poisoned) BeforeAction() {
	s.role.SubHp(30)
}

func (s *Poisoned) AfterAction() {
	s.round -= 1
}

func (s *Poisoned) IsFinished() bool {
	return s.round <= 0
}

func (s *Poisoned) MakeDamage(damage int) int {
	return damage
}
