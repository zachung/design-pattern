package skill

import (
	"3B/internal/contract"
	"fmt"
	"log"
	"strings"
)

type Waterball struct {
	name   string
	mpCost int
	damage int
}

func NewWaterball() contract.Skill {
	return &Waterball{
		name:   "水球",
		mpCost: 50,
		damage: 120,
	}
}

func (a *Waterball) GetName() string {
	return a.name
}

func (a *Waterball) CanCast(role contract.Role) bool {
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *Waterball) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
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
		damage := role.MakeDamage(a.damage)
		log.Println(fmt.Sprintf("%s 對 %s 造成 %d 點傷害。", role.GetName(), enemy.GetName(), damage))
		enemy.Property(contract.Hp).Sub(damage)
	}
}
