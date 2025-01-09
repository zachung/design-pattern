package internal

import (
	"3B/internal/skill"
	"fmt"
	"strings"
)

type Round struct{}

func RunRound(role *Role, targetTroop *Troop) {
	var s *skill.Skill
	for {
		s = selectSkill(role)
		if s != nil {
			break
		}
	}
	targetCount := s.Area[skill.Enemy]
	if targetCount != 0 {
		targets := targetTroop.AliveRoles()
		if targetCount > 0 && targetCount < len(targets) {
			// 需要選擇敵方目標
			var targetList []string
			for i, t := range targets {
				targetList = append(targetList, fmt.Sprintf("(%d) %s", i, t.GetName()))
			}
			fmt.Printf("選擇 %d 位目標: %s\n", targetCount, strings.Join(targetList, " "))
			targets = selectTarget(role, targets)
		}
		castSkill(role, s, targets)
	}
}

func selectSkill(role *Role) *skill.Skill {
	var actionList []string
	for i, s := range role.Skills {
		actionList = append(actionList, fmt.Sprintf("(%d) %s", i, s))
	}
	// print 選擇行動
	fmt.Printf("選擇行動：%v\n", strings.Join(actionList, " "))

	command := role.controller.PullCommand()

	return getSkill(role, command[0])
}

func getSkill(role *Role, i int) (s *skill.Skill) {
	str := role.Skills[i]
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
	if role.Mp < s.MpCost {
		fmt.Println("你缺乏 MP，不能進行此行動。")
		return nil
	}
	return s
}

func selectTarget(role *Role, enemies []*Role) (targets []*Role) {
	// 選擇敵方目標
	command := role.controller.PullCommand()
	for _, i := range command {
		targets = append(targets, enemies[i])
	}

	return
}

func castSkill(role *Role, s *skill.Skill, targets []*Role) {
	role.Mp -= s.MpCost
	var str []string
	for _, enemy := range targets {
		str = append(str, enemy.GetName())
	}
	if s == &skill.BasicAttack {
		fmt.Printf("%s 攻擊 %s。\n", role.GetName(), strings.Join(str, ", "))
	} else {
		fmt.Printf("%s 對 %s 使用了 %s。\n", role.GetName(), strings.Join(str, ", "), s.Name)
	}
	for _, enemy := range targets {
		damage := s.Damage
		if s == &skill.BasicAttack {
			damage = role.Str
		}
		enemy.Hp -= damage
		fmt.Printf("%s 對 %s 造成 %d 點傷害。\n", role.GetName(), enemy.GetName(), damage)
		if enemy.IsDead() {
			fmt.Printf("%s 死亡。\n", enemy.GetName())
		}
	}
}
