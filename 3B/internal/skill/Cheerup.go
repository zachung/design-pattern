package skill

import (
	"3B/internal/contract"
	"3B/internal/state"
	"fmt"
	"log"
	"strings"
)

type Cheerup struct {
	name   string
	mpCost int
}

func NewCheerup() contract.Skill {
	return &Cheerup{
		name:   "鼓舞",
		mpCost: 100,
	}
}

func (a *Cheerup) GetName() string {
	return a.name
}

func (a *Cheerup) CanCast(role contract.Role) bool {
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *Cheerup) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	targetCount := 3
	targets := ally.AliveRoles()
	for i, target := range targets {
		if target == role {
			// 除去自己
			targets = append(targets[:i], targets[i+1:]...)
			break
		}
	}
	if targetCount < len(targets) {
		targets = role.Actor().SelectTarget(targets, targetCount)
	}
	role.Property(contract.Mp).Sub(a.mpCost)

	var str []string
	for _, enemy := range targets {
		str = append(str, enemy.GetName())
	}
	join := strings.Join(str, ", ")
	if join != "" {
		join = fmt.Sprintf("對 %s ", join)
	}
	log.Println(fmt.Sprintf("%s %s使用了 %s。", role.GetName(), join, a.GetName()))
	for _, enemy := range targets {
		enemy.SetState(state.NewCheerup(enemy))
	}
}
