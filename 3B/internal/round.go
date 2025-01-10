package internal

import (
	"3B/internal/skill"
)

type Actor interface {
	SelectSkill() *skill.Skill
	SelectTarget(enemies []*Role, targetCount int) []*Role
	CastSkill(s *skill.Skill, targets []*Role)
}

type Round struct {
	role *Role
}

func (r *Round) SelectSkill() *skill.Skill {
	selected := r.role.controller.PullCommand()[0]
	return r.role.SelectSkill(selected)
}

func (r *Round) SelectTarget(enemies []*Role, targetCount int) []*Role {
	command := r.role.controller.PullCommand()
	return r.role.SelectTarget(enemies, targetCount, command)
}

func (r *Round) CastSkill(s *skill.Skill, targets []*Role) {
	r.role.CastSkill(s, targets)
}
