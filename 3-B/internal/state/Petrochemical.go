package state

type Petrochemical struct {
	name  string
	round int
}

func NewPetrochemical() *Petrochemical {
	return &Petrochemical{name: "石化", round: 3}
}

func (s *Petrochemical) GetName() string {
	return s.name
}

func (s *Petrochemical) CanAction() bool {
	return false
}

func (s *Petrochemical) BeforeAction() {
}

func (s *Petrochemical) AfterAction() {
	s.round -= 1
}

func (s *Petrochemical) IsFinished() bool {
	return s.round <= 0
}

func (s *Petrochemical) MakeDamage(damage int) int {
	return damage
}
