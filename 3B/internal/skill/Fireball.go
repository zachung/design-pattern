package skill

import (
	"3B/internal/contract"
	"fmt"
	"log"
	"strings"
)

type Fireball struct {
	name   string
	mpCost int
	damage int
}

func NewFireball() contract.Skill {
	return &Fireball{
		name:   "火球",
		mpCost: 50,
		damage: 50,
	}
}

func (a *Fireball) GetName() string {
	return a.name
}

func (a *Fireball) CanCast(role contract.Role) bool {
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *Fireball) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	targets := enemy.AliveRoles()

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
