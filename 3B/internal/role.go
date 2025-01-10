package internal

import (
	"3B/internal/skill"
	"fmt"
	"strings"
)

type Role struct {
	Name        string
	Hp, Mp, Str int
	Skills      []string
	State       string
	troop       *Troop
	actor       Actor
	commands    []string
	controller  *Controller
}

func NewRole(name string, hp, mp, str int, skills []string) *Role {
	role := &Role{
		Name: name, Hp: hp, Mp: mp, Str: str, Skills: skills,
		State:      "正常",
		controller: NewController(),
	}
	var actor Actor
	if name == "英雄" {
		actor = &Round{role: role}
	} else {
		actor = &AI{role: role}
	}
	role.actor = actor

	return role
}

func (r *Role) IsDead() bool {
	return r.Hp <= 0
}

func (r *Role) GetName() string {
	return fmt.Sprintf("[%d]%s", r.troop.I, r.Name)
}

func (r *Role) Action(targetTroop *Troop) {
	fmt.Printf("輪到 %s (HP: %d, MP: %d, STR: %d, State: %s)。\n", r.GetName(), r.Hp, r.Mp, r.Str, r.State)
	r.actor.Action(targetTroop)
}

func (r *Role) SelectSkill(selected int) *skill.Skill {
	var actionList []string
	for i, s := range r.Skills {
		actionList = append(actionList, fmt.Sprintf("(%d) %s", i, s))
	}
	// print 選擇行動
	fmt.Printf("選擇行動：%v\n", strings.Join(actionList, " "))

	var s *skill.Skill
	str := r.Skills[selected]
	switch str {
	case "普通攻擊":
		return &skill.BasicAttack
	case "水球":
		s = &skill.Waterball
	case "火球":
		s = &skill.Fireball
	}
	if s == nil {
		return nil
	}
	if r.Mp < s.MpCost {
		fmt.Println("你缺乏 MP，不能進行此行動。")
		return nil
	}
	return s
}

func (r *Role) SelectTarget(enemies []*Role, targetCount int, selected []int) (targets []*Role) {
	// 需要選擇敵方目標
	var targetList []string
	for i, t := range enemies {
		targetList = append(targetList, fmt.Sprintf("(%d) %s", i, t.GetName()))
	}
	fmt.Printf("選擇 %d 位目標: %s\n", targetCount, strings.Join(targetList, " "))

	// 選擇敵方目標
	for _, i := range selected {
		targets = append(targets, enemies[i])
	}

	return
}

func (r *Role) CastSkill(s *skill.Skill, targets []*Role) {
	r.Mp -= s.MpCost
	var str []string
	for _, enemy := range targets {
		str = append(str, enemy.GetName())
	}
	if s == &skill.BasicAttack {
		fmt.Printf("%s 攻擊 %s。\n", r.GetName(), strings.Join(str, ", "))
	} else {
		fmt.Printf("%s 對 %s 使用了 %s。\n", r.GetName(), strings.Join(str, ", "), s.Name)
	}
	for _, enemy := range targets {
		damage := s.Damage
		if s == &skill.BasicAttack {
			damage = r.Str
		}
		enemy.Hp -= damage
		fmt.Printf("%s 對 %s 造成 %d 點傷害。\n", r.GetName(), enemy.GetName(), damage)
		if enemy.IsDead() {
			fmt.Printf("%s 死亡。\n", enemy.GetName())
		}
	}
}
