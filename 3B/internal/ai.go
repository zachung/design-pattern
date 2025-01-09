package internal

import (
	"3B/internal/skill"
	"fmt"
	"strconv"
	"strings"
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
	a.role.controller.AddCommand(strconv.Itoa(a.seed % len(a.role.Skills)))
	return a.role.SelectSkill()
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
	cmd := strings.Trim(strings.Join(strings.Split(fmt.Sprint(ints), " "), ", "), "[]")
	a.role.controller.AddCommand(cmd)

	return a.role.SelectTarget(enemies, targetCount)
}
