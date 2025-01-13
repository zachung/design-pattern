package skill

import (
	"3B/internal/contract"
	"3B/internal/state"
	"fmt"
	"log"
	"strings"
)

type Poison struct {
	name   string
	mpCost int
}

func NewPoison() contract.Skill {
	return &Poison{
		name:   "下毒",
		mpCost: 80,
	}
}

func (a *Poison) GetName() string {
	return a.name
}

func (a *Poison) CanCast(role contract.Role) bool {
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *Poison) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
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
	log.Println(fmt.Sprintf("%s 對 %s 使用了 %s。", role.GetName(), strings.Join(str, ", "), a.GetName()))
	for _, enemy := range targets {
		enemy.SetState(state.NewPoisoned(enemy))
	}
}
