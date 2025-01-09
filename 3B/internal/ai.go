package internal

import (
	"3B/internal/skill"
	"fmt"
	"strconv"
	"strings"
)

type AI struct {
	seed int
}

func (a *AI) Action(role *Role, targetTroop *Troop) {
	defer func() {
		a.seed += 1
	}()
	s := a.selectSkill(role)
	targetCount := s.Area[skill.Enemy]
	if targetCount != 0 {
		targets := targetTroop.AliveRoles()
		if targetCount > 0 && targetCount < len(targets) {
			a.selectTarget(role, len(targets))
			// 選擇敵方目標
			targets = selectTarget(role, targets)
		}
		castSkill(role, s, targets)
	}
}

func (a *AI) selectSkill(role *Role) *skill.Skill {
	var s *skill.Skill
	actionCount := len(role.Skills)
	for {
		role.controller.AddCommand(strconv.Itoa(a.seed % actionCount))
		s = selectSkill(role)
		if s != nil {
			break
		}
		a.seed += 1
	}
	return s
}

func (a *AI) selectTarget(role *Role, n int) {
	var ints []int
	for i := a.seed % n; i < n; i++ {
		if i == n {
			i = 0
		}
		ints = append(ints, i)
	}
	cmd := strings.Trim(strings.Join(strings.Split(fmt.Sprint(ints), " "), ", "), "[]")
	role.controller.AddCommand(cmd)
}
