package state

type NormalState struct {
	name string
}

func NewNormalState() *NormalState {
	return &NormalState{name: "正常"}
}

func (s *NormalState) GetName() string {
	return s.name
}

func (s *NormalState) CanAction() bool {
	return true
}

func (s *NormalState) BeforeAction() {
}

func (s *NormalState) AfterAction() {
}

func (s *NormalState) IsFinished() bool {
	return false
}
