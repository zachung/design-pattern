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
	case "下毒":
		return NewPoison()
	case "召喚":
		return NewSummon()
	case "自爆":
		return NewSelfExplosion()
	case "鼓舞":
		return NewCheerup()
	case "詛咒":
		return NewCurse()
	case "一拳攻擊":
		return NewOnePunch()
	}
	panic("skill " + name + " not found")
}
