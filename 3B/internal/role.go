package internal

import "fmt"

type Role struct {
	Name        string
	Hp, Mp, Str int
	Skills      []string
	State       string
	troop       *Troop
	ai          AI
	commands    []string
	controller  *Controller
}

func NewRole(name string, hp, mp, str int, skills []string) *Role {
	ai := AI{}
	return &Role{
		Name: name, Hp: hp, Mp: mp, Str: str, Skills: skills,
		State:      "正常",
		ai:         ai,
		controller: NewController(),
	}
}

func (r *Role) IsDead() bool {
	return r.Hp <= 0
}

func (r *Role) GetName() string {
	return fmt.Sprintf("[%d]%s", r.troop.I, r.Name)
}

func (r *Role) Action(targetTroop *Troop) {
	fmt.Printf("輪到 %s (HP: %d, MP: %d, STR: %d, State: %s)。\n", r.GetName(), r.Hp, r.Mp, r.Str, r.State)
	if r.Name == "英雄" {
		RunRound(r, targetTroop)
	} else {
		r.ai.Action(r, targetTroop)
	}
}
