package internal

import (
	"3B/internal/contract"
)

type AI struct {
	role contract.Role
	seed int
}

func (a *AI) SelectSkill(skillCount int) contract.Skill {
	s := a.role.SelectSkill(a.seed % skillCount)
	a.seed += 1
	return s
}

func (a *AI) SelectTarget(enemies []contract.Role, targetCount int) []contract.Role {
	n := len(enemies)
	var ints []int
	for i := a.seed % n; i < n; i++ {
		if i == n {
			i = 0
		}
		ints = append(ints, i)
		if len(ints) == targetCount {
			// 已選到足夠數量
			break
		}
	}

	return a.role.SelectTarget(enemies, targetCount, ints)
}

func (a *AI) CastSkill(s contract.Skill, ally contract.Troop, enemy contract.Troop) {
	s.Cast(a.role, ally, enemy)
	a.seed += 1
}
