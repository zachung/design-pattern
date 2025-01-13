package skill

import (
	"3B/internal/contract"
	"fmt"
	"strings"
)

type SelfExplosion struct {
	name   string
	mpCost int
	damage int
}

func NewSelfExplosion() contract.Skill {
	return &SelfExplosion{
		name:   "自爆",
		mpCost: 200,
		damage: 150,
	}
}

func (a *SelfExplosion) GetName() string {
	return a.name
}

func (a *SelfExplosion) CanCast(role contract.Role) bool {
	return role.GetMp() >= a.mpCost
}

func (a *SelfExplosion) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	targets := ally.AliveRoles()
	for i, target := range targets {
		if target == role {
			// 除去自己
			targets = append(targets[:i], targets[i+1:]...)
			break
		}
	}
	targets = append(targets, enemy.AliveRoles()...)

	role.SubMp(a.mpCost)

	var str []string
	for _, enemy := range targets {
		str = append(str, enemy.GetName())
	}
	fmt.Printf("%s 對 %s 使用了 %s。\n", role.GetName(), strings.Join(str, ", "), a.GetName())
	for _, enemy := range targets {
		damage := role.MakeDamage(a.damage)
		fmt.Printf("%s 對 %s 造成 %d 點傷害。\n", role.GetName(), enemy.GetName(), damage)
		enemy.SubHp(damage)
	}
	// 使自己死亡
	role.SubHp(role.GetHp())
}
