package internal

import (
	"3B/internal/contract"
)

type Round struct {
	role       contract.Role
	controller *Controller
}

func (r *Round) SelectSkill(skillCount int) contract.Skill {
	selected := r.controller.PullCommand()[0]
	return r.role.SelectSkill(selected)
}

func (r *Round) SelectTarget(enemies []contract.Role, targetCount int) []contract.Role {
	command := r.controller.PullCommand()
	return r.role.SelectTarget(enemies, targetCount, command)
}

func (r *Round) CastSkill(s contract.Skill, ally contract.Troop, enemy contract.Troop) {
	s.Cast(r.role, ally, enemy)
}
