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
		selected := r.role.controller.PullCommand()[0]
		s = r.role.SelectSkill(selected)
		if s != nil {
			break
		}
	}
	targetCount := s.Area[skill.Enemy]
	if targetCount != 0 {
		targets := targetTroop.AliveRoles()
		if targetCount > 0 && targetCount < len(targets) {
			command := r.role.controller.PullCommand()
			targets = r.role.SelectTarget(targets, targetCount, command)
		}
		r.role.CastSkill(s, targets)
	}
}
