package skill

import (
	"3-B/internal/actor"
	"3-B/internal/contract"
	"fmt"
	"log"
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
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *Summon) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	role.Property(contract.Mp).Sub(a.mpCost)
	log.Println(fmt.Sprintf("%s 使用了 %s。", role.GetName(), a.GetName()))
	slime := ally.NewRole("Slime 100 0 50")
	slime.Property(contract.Hp).AddObserver(func(v *int) {
		if *v <= 0 {
			role.Property(contract.Hp).Add(30)
		}
	})
	ally.AddRole(slime)
	slime.SetActor(actor.NewAI(slime))
}
