package contract

type CardPattern interface {
	IsGreaterThan(CardPattern) bool
	String() string
	GetCards() []Card
}

type Pattern struct {
	CardPattern
	Cards []Card
}

func (p *Pattern) GetCards() []Card {
	return p.Cards
}
