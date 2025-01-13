package skill

import (
	"3B/internal/contract"
	"fmt"
)

type Summon struct {
	name   string
	mpCost int
	addHp  int
}

func NewSummon() contract.Skill {
	return &Summon{
		name:   "召喚",
		mpCost: 150,
	}
}

func (a *Summon) GetName() string {
	return a.name
}

func (a *Summon) CanCast(role contract.Role) bool {
	return role.GetMp() >= a.mpCost
}

func (a *Summon) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	role.SubMp(a.mpCost)
	role.AddHp(a.addHp)
	fmt.Printf("%s 使用了 %s。\n", role.GetName(), a.GetName())
	slime := ally.NewRole("Slime 100 0 50")
	slime.SetObserver(contract.OnDead, func() {
		role.AddHp(30)
	})
	ally.AddRole(slime)
}
