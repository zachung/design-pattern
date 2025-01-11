package skill

import (
	"3B/internal/contract"
	"fmt"
)

type SelfHealing struct {
	Name string
}

func (a *SelfHealing) GetName() string {
	return "自我治療"
}

func (a *SelfHealing) CanCast(role contract.Role) bool {
	return role.GetMp() >= 50
}

func (a *SelfHealing) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	role.SubMp(50)
	role.AddHp(150)
	fmt.Printf("%s 使用了 %s。\n", role.GetName(), a.GetName())
}
