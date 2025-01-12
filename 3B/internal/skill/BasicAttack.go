package skill

import (
	"3B/internal/contract"
	"fmt"
	"strings"
)

type BasicAttack struct {
	name string
}

func NewBasicAttack() contract.Skill {
	return &BasicAttack{
		name: "普通攻擊",
	}
}

func (a *BasicAttack) GetName() string {
	return a.name
}

func (a *BasicAttack) CanCast(role contract.Role) bool {
	return true
}

func (a *BasicAttack) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	targetCount := 1
	targets := enemy.AliveRoles()
	if targetCount < len(targets) {
		targets = role.Actor().SelectTarget(targets, targetCount)
	}

	var str []string
	for _, enemy := range targets {
		str = append(str, enemy.GetName())
	}
	fmt.Printf("%s 攻擊 %s。\n", role.GetName(), strings.Join(str, ", "))
	for _, enemy := range targets {
		damage := role.GetStr()
		fmt.Printf("%s 對 %s 造成 %d 點傷害。\n", role.GetName(), enemy.GetName(), damage)
		enemy.SubHp(damage)
	}
}
