package internal

import (
	"3B/internal/skill"
)

type AI struct {
	role *Role
	seed int
}

func (a *AI) SelectSkill() *skill.Skill {
	s := a.role.SelectSkill(a.seed % len(a.role.Skills))
	if s == nil {
		a.seed += 1
	}
	return s
}

func (a *AI) SelectTarget(enemies []*Role, targetCount int) []*Role {
	n := len(enemies)
	var ints []int
	for i := a.seed % n; i < n; i++ {
		if i == n {
			i = 0
		}
		ints = append(ints, i)
	}

	return a.role.SelectTarget(enemies, targetCount, ints)
}

func (a *AI) CastSkill(s *skill.Skill, targets []*Role) {
	a.role.CastSkill(s, targets)
	a.seed += 1
}
