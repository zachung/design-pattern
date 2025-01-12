package skill

import "3B/internal/contract"

func GetSkill(name string) contract.Skill {
	switch name {
	case "普通攻擊":
		return NewBasicAttack()
	case "水球":
		return NewWaterball()
	case "火球":
		return NewFireball()
	case "自我治療":
		return NewSelfHealing()
	case "石化":
		return NewPetrochemical()
	}
	panic("skill " + name + " not found")
}
