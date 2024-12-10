package objects

import (
	"3-3/internal/contract"
	"3-3/internal/states"
	"math/rand/v2"
)

type TreasureName string

const (
	superStar          TreasureName = "無敵星星"
	poison             TreasureName = "毒藥"
	acceleratingPotion TreasureName = "加速藥水"
	healingPotion      TreasureName = "補血罐"
	devilFruit         TreasureName = "惡魔果實"
	kingsRock          TreasureName = "王者之印"
	dokodemoDoor       TreasureName = "任意門"
)

type Treasure struct {
	GameObject
	treasureName TreasureName
}

func NewTreasure(gameMap *contract.Map) *Treasure {
	return &Treasure{
		GameObject: GameObject{
			gameMap: gameMap,
			name:    contract.TreasureName,
		},
		treasureName: randState(),
	}
}

func randState() TreasureName {
	n := rand.IntN(100)
	if n < 10 {
		return superStar
	}
	if n < 35 {
		return poison
	}
	if n < 55 {
		return acceleratingPotion
	}
	if n < 70 {
		return healingPotion
	}
	if n < 80 {
		return devilFruit
	}
	if n < 90 {
		return kingsRock
	}
	return dokodemoDoor
}

func (o *Treasure) Symbol() rune {
	return contract.TreasureSymbol
}

func (o *Treasure) GetTreasureName() string {
	return string(o.treasureName)
}

func (o *Treasure) GetState(role *contract.Role) *contract.State {
	switch o.treasureName {
	case superStar:
		return states.NewInvincible(role)
	case poison:
		return states.NewPoisoned(role)
	case acceleratingPotion:
		return states.NewAccelerated(role)
	case healingPotion:
		return states.NewHealing(role)
	case devilFruit:
		return states.NewOrderless(role)
	case kingsRock:
		return states.NewStockpile(role)
	case dokodemoDoor:
		return states.NewTeleport(role)
	}
	return states.NewNormal(role)
}
