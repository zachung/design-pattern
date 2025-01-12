package skill

import (
	"3B/internal/contract"
	"fmt"
)

type SelfHealing struct {
	name   string
	mpCost int
	addHp  int
}

func NewSelfHealing() contract.Skill {
	return &SelfHealing{
		name:   "自我治療",
		mpCost: 50,
		addHp:  150,
	}
}

func (a *SelfHealing) GetName() string {
	return a.name
}

func (a *SelfHealing) CanCast(role contract.Role) bool {
	return role.GetMp() >= a.mpCost
}

func (a *SelfHealing) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	role.SubMp(a.mpCost)
	role.AddHp(a.addHp)
	fmt.Printf("%s 使用了 %s。\n", role.GetName(), a.GetName())
}
