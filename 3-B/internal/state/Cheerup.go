package state

import "3-B/internal/contract"

type Cheerup struct {
	name  string
	round int
	role  contract.Role
}

func NewCheerup(role contract.Role) *Cheerup {
	return &Cheerup{name: "受到鼓舞", round: 3, role: role}
}

func (s *Cheerup) GetName() string {
	return s.name
}

func (s *Cheerup) CanAction() bool {
	return true
}

func (s *Cheerup) BeforeAction() {
}

func (s *Cheerup) AfterAction() {
	s.round -= 1
}

func (s *Cheerup) IsFinished() bool {
	return s.round <= 0
}

func (s *Cheerup) MakeDamage(damage int) int {
	return damage + 50
}
