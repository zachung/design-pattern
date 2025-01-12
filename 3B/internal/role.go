package internal

import (
	"3B/internal/contract"
	"3B/internal/skill"
	"3B/internal/state"
	"fmt"
	"strings"
)

type RoleImpl struct {
	Name        string
	Hp, Mp, Str int
	MaxHp       int
	Skills      []string
	State       contract.State
	troop       contract.Troop
	actor       contract.Actor
	commands    []string
	observers   map[contract.Event]func()
}

func NewRole(troop contract.Troop, name string, hp, mp, str int, skills []string) contract.Role {
	return &RoleImpl{
		Name:      name,
		Hp:        hp,
		MaxHp:     hp,
		Mp:        mp,
		Str:       str,
		Skills:    skills,
		State:     state.GetState("正常"),
		troop:     troop,
		observers: map[contract.Event]func(){},
	}
}

func (r *RoleImpl) SetActor(actor contract.Actor) {
	r.actor = actor
}

func (r *RoleImpl) IsDead() bool {
	return r.Hp <= 0
}

func (r *RoleImpl) GetName() string {
	return fmt.Sprintf("[%d]%s", r.troop.GetI(), r.Name)
}

func (r *RoleImpl) Action(targetTroop contract.Troop) {
	defer func() {
		r.State.AfterAction()
		if r.State.IsFinished() {
			r.SetState(state.GetState("正常"))
		}
	}()
	fmt.Printf("輪到 %s (HP: %d, MP: %d, STR: %d, State: %s)。\n", r.GetName(), r.Hp, r.Mp, r.Str, r.State.GetName())
	if !r.State.CanAction() {
		return
	}
	r.State.BeforeAction()
	if r.IsDead() {
		return
	}
	var s contract.Skill
	for {
		s = r.actor.SelectSkill(len(r.Skills))
		if s != nil {
			break
		}
	}
	r.actor.CastSkill(s, r.troop, targetTroop)
}

func (r *RoleImpl) SelectSkill(selected int) contract.Skill {
	var actionList []string
	for i, s := range r.Skills {
		actionList = append(actionList, fmt.Sprintf("(%d) %s", i, s))
	}
	// print 選擇行動
	fmt.Printf("選擇行動：%v\n", strings.Join(actionList, " "))

	s := skill.GetSkill(r.Skills[selected])
	if s == nil {
		return nil
	}
	if !s.CanCast(r) {
		fmt.Println("你缺乏 MP，不能進行此行動。")
		return nil
	}
	return s
}

func (r *RoleImpl) SubHp(damage int) {
	r.Hp -= damage
	if r.IsDead() {
		f := r.observers[contract.OnDead]
		if f != nil {
			f()
		}
		fmt.Printf("%s 死亡。\n", r.GetName())
	}
}

func (r *RoleImpl) AddHp(health int) {
	r.Hp += health
	if r.Hp > r.MaxHp {
		r.Hp = r.MaxHp
	}
}

func (r *RoleImpl) GetStr() int {
	return r.Str
}

func (r *RoleImpl) GetMp() int {
	return r.Mp
}

func (r *RoleImpl) SubMp(mp int) {
	r.Mp -= mp
}

func (r *RoleImpl) SetState(state contract.State) {
	r.State = state
}

func (r *RoleImpl) Actor() contract.Actor {
	return r.actor
}

func (r *RoleImpl) SelectTarget(enemies []contract.Role, targetCount int, selected []int) (targets []contract.Role) {
	// 選擇敵方目標
	for _, i := range selected {
		targets = append(targets, enemies[i])
	}

	return
}

func (r *RoleImpl) SetObserver(event contract.Event, observer func()) {
	r.observers[event] = observer
}
