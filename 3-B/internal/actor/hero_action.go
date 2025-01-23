package actor

import (
	"3-B/internal/contract"
	"fmt"
	"log"
	"strings"
)

type HeroAction struct {
	role       contract.Role
	controller *Controller
}

func NewHeroAction(role contract.Role, controller *Controller) contract.Actor {
	return &HeroAction{role, controller}
}

func (r *HeroAction) SelectSkill(skillCount int) contract.Skill {
	selected := r.controller.PullCommand()[0]
	return r.role.SelectSkill(selected)
}

func (r *HeroAction) SelectTarget(enemies []contract.Role, targetCount int) []contract.Role {
	// 需要選擇敵方目標
	var targetList []string
	for i, t := range enemies {
		targetList = append(targetList, fmt.Sprintf("(%d) %s", i, t.GetName()))
	}
	log.Println(fmt.Sprintf("選擇 %d 位目標: %s", targetCount, strings.Join(targetList, " ")))
	command := r.controller.PullCommand()
	return r.role.SelectTarget(enemies, targetCount, command)
}

func (r *HeroAction) CastSkill(s contract.Skill, ally contract.Troop, enemy contract.Troop) {
	s.Cast(r.role, ally, enemy)
}
