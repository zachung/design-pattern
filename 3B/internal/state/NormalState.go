package state

type NormalState struct{}

func NewNormalState() *NormalState {
	return &NormalState{}
}

func (s *NormalState) BeforeRound() {

}

func (s *NormalState) AfterRound() {
}

func (s *NormalState) IsFinished() bool {
	return false
}
