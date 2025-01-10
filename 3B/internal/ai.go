package internal

import (
	"3B/internal/skill"
)

type AI struct {
	role *Role
	seed int
}

func (a *AI) Action(targetTroop *Troop) {
	var s *skill.Skill
	for {
		s = a.selectSkill()
		if s != nil {
			break
		}
		a.seed += 1
	}
	targetCount := s.Area[skill.Enemy]
	if targetCount != 0 {
		targets := targetTroop.AliveRoles()
		if targetCount > 0 && targetCount < len(targets) {
			targets = a.selectTarget(targets, targetCount)
		}
		a.role.CastSkill(s, targets)
	}
	a.seed += 1
}

func (a *AI) selectSkill() *skill.Skill {
	return a.role.SelectSkill(a.seed % len(a.role.Skills))
}

func (a *AI) selectTarget(enemies []*Role, targetCount int) []*Role {
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
