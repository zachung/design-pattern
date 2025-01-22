package internal

import (
	"3B/internal/contract"
	"3B/internal/skill"
	"3B/internal/state"
	"fmt"
	"log"
	"strings"
)

type RoleImpl struct {
	Name       string
	properties map[contract.PropertyKey]contract.Property
	MaxHp      int
	Skills     []contract.Skill
	State      contract.State
	troop      contract.Troop
	actor      contract.Actor
	commands   []string
}

func NewRole(troop contract.Troop, name string, hp, mp, str int, skills []string) contract.Role {
	properties := map[contract.PropertyKey]contract.Property{
		contract.Hp:  NewProperty(hp),
		contract.Mp:  NewProperty(mp),
		contract.Str: NewProperty(str),
	}
	role := &RoleImpl{
		Name:       name,
		MaxHp:      hp,
		Skills:     initSkills(skills),
		State:      state.NewNormalState(),
		troop:      troop,
		properties: properties,
	}
	// observer dead event
	properties[contract.Hp].AddObserver(func(v *int) {
		if *v <= 0 {
			log.Println(fmt.Sprintf("%s 死亡。", role.GetName()))
		}
	})
	return role
}

func initSkills(skillNames []string) (skills []contract.Skill) {
	for _, skillName := range skillNames {
		skills = append(skills, skill.GetSkill(skillName))
	}
	return
}

func (r *RoleImpl) SetActor(actor contract.Actor) {
	r.actor = actor
}

func (r *RoleImpl) IsDead() bool {
	return r.Property(contract.Hp).Get() <= 0
}

func (r *RoleImpl) GetName() string {
	return fmt.Sprintf("[%d]%s", r.troop.GetI(), r.Name)
}

func (r *RoleImpl) Action(targetTroop contract.Troop) {
	defer func() {
		r.State.AfterAction()
		if r.State.IsFinished() {
			r.SetState(state.NewNormalState())
		}
	}()
	turnMessage := fmt.Sprintf(
		"輪到 %s (HP: %d, MP: %d, STR: %d, State: %s)。",
		r.GetName(),
		r.Property(contract.Hp).Get(),
		r.Property(contract.Mp).Get(),
		r.Property(contract.Str).Get(),
		r.State.GetName(),
	)
	log.Println(turnMessage)
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
		actionList = append(actionList, fmt.Sprintf("(%d) %s", i, s.GetName()))
	}
	// print 選擇行動
	log.Println(fmt.Sprintf("選擇行動：%v", strings.Join(actionList, " ")))

	s := r.Skills[selected]
	if s == nil {
		return nil
	}
	if !s.CanCast(r) {
		log.Println("你缺乏 MP，不能進行此行動。")
		return nil
	}
	return s
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

func (r *RoleImpl) MakeDamage(damage int) int {
	return r.State.MakeDamage(damage)
}

func (r *RoleImpl) Property(key contract.PropertyKey) contract.Property {
	return r.properties[key]
}

func (r *RoleImpl) GetState() contract.State {
	return r.State
}
