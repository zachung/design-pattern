package skill

import (
	"3B/internal/contract"
	"3B/internal/state"
	"fmt"
	"strings"
)

type Petrochemical struct {
	name   string
	mpCost int
}

func NewPetrochemical() contract.Skill {
	return &Petrochemical{
		name:   "石化",
		mpCost: 100,
	}
}

func (a *Petrochemical) GetName() string {
	return a.name
}

func (a *Petrochemical) CanCast(role contract.Role) bool {
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *Petrochemical) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	targetCount := 1
	targets := enemy.AliveRoles()
	if targetCount < len(targets) {
		targets = role.Actor().SelectTarget(targets, targetCount)
	}
	role.Property(contract.Mp).Sub(a.mpCost)

	var str []string
	for _, enemy := range targets {
		str = append(str, enemy.GetName())
	}
	fmt.Printf("%s 對 %s 使用了 %s。\n", role.GetName(), strings.Join(str, ", "), a.GetName())
	for _, enemy := range targets {
		enemy.SetState(state.NewPetrochemical())
	}
}
