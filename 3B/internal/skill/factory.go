package skill

import "3B/internal/contract"

func GetSkill(name string) contract.Skill {
	switch name {
	case "普通攻擊":
		return &BasicAttack{}
	case "水球":
		return &Waterball{}
	case "火球":
		return &Fireball{}
	case "自我治療":
		return &SelfHealing{}
	}
	panic("skill " + name + " not found")
}
