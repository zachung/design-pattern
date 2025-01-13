package skill

import (
	"3B/internal/contract"
	"fmt"
	"log"
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
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *SelfHealing) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	role.Property(contract.Mp).Sub(a.mpCost)
	role.Property(contract.Hp).Add(a.addHp)
	log.Println(fmt.Sprintf("%s 使用了 %s。", role.GetName(), a.GetName()))
}
