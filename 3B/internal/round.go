package internal

import (
	"3B/internal/skill"
)

type Actor interface {
	Action(targetTroop *Troop)
}

type Round struct {
	role *Role
}

func (r *Round) Action(targetTroop *Troop) {
	var s *skill.Skill
	for {
		s = r.role.SelectSkill()
		if s != nil {
			break
		}
	}
	targetCount := s.Area[skill.Enemy]
	if targetCount != 0 {
		targets := targetTroop.AliveRoles()
		if targetCount > 0 && targetCount < len(targets) {
			targets = r.role.SelectTarget(targets, targetCount)
		}
		r.role.CastSkill(s, targets)
	}
}
