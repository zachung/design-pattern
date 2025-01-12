package skill

import (
	"3B/internal/contract"
	"fmt"
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
	return role.GetMp() >= a.mpCost
}

func (a *Fireball) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	targets := enemy.AliveRoles()

	role.SubMp(a.mpCost)

	var str []string
	for _, enemy := range targets {
		str = append(str, enemy.GetName())
	}
	fmt.Printf("%s 對 %s 使用了 %s。\n", role.GetName(), strings.Join(str, ", "), a.GetName())
	for _, enemy := range targets {
		damage := a.damage
		enemy.SubHp(damage)
		fmt.Printf("%s 對 %s 造成 %d 點傷害。\n", role.GetName(), enemy.GetName(), damage)
		if enemy.IsDead() {
			fmt.Printf("%s 死亡。\n", enemy.GetName())
		}
	}
}
