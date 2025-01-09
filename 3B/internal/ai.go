package internal

import (
	skill2 "3B/internal/skill"
	"fmt"
	"strings"
)

type AI struct {
	seed int
}

func (a *AI) Action(role *Role, targetTroop *Troop) {
	defer func() {
		a.seed += 1
	}()
	actionCount := len(role.Skills)
	s := a.seed % actionCount
	skill := role.Skills[s]
	switch skill {
	case "普通攻擊":
		enemies := targetTroop.AliveRoles()
		n := len(enemies)
		// 只需選擇一個
		i := a.seed % n
		target := enemies[i]
		target.Hp -= role.Str
		fmt.Printf("%s 對 %s 使用了 %s。\n", role.GetName(), target.GetName(), skill)
		fmt.Printf("%s 對 %s 造成 %d 點傷害。\n", role.GetName(), target.GetName(), role.Str)
		return
	case "水球":
		enemies := targetTroop.AliveRoles()
		n := len(enemies)
		// 只需選擇一個
		i := a.seed % n
		target := enemies[i]
		target.Hp -= skill2.Waterball.Damage
		fmt.Printf("%s 對 %s 使用了 %s。\n", role.GetName(), target.GetName(), skill)
		fmt.Printf("%s 對 %s 造成 %d 點傷害。\n", role.GetName(), target.GetName(), skill2.Waterball.Damage)
		return
	case "火球":
		enemies := targetTroop.AliveRoles()
		var str []string
		for _, enemy := range enemies {
			str = append(str, enemy.GetName())
		}
		fmt.Printf("%s 對 %s 使用了 %s。\n", role.GetName(), strings.Join(str, ", "), skill)
		for _, enemy := range enemies {
			enemy.Hp -= skill2.Fireball.Damage
			fmt.Printf("%s 對 %s 造成 %d 點傷害。\n", role.GetName(), enemy.GetName(), skill2.Fireball.Damage)
		}
		return
	}
}
