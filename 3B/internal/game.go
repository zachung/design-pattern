package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	battle Battle
}

func Run(initConfig []string) {
	var troop *Troop
	var battle = Battle{troops: make([]*Troop, 0)}
	for _, s := range initConfig {
		r := regexp.MustCompile(`#軍隊-(\d)-(.*)`)
		match := r.FindStringSubmatch(s)
		if len(match) > 1 {
			i, _ := strconv.Atoi(match[1])
			// 切換軍隊
			if match[2] == "開始" {
				troop = &Troop{I: i}
			}
			if match[2] == "結束" {
				battle.troops = append(battle.troops, troop)
			}
			continue
		}
		if len(battle.troops) < 2 {
			// 添加角色
			properties := strings.Split(s, " ")
			hp, _ := strconv.Atoi(properties[1])
			mp, _ := strconv.Atoi(properties[2])
			str, _ := strconv.Atoi(properties[3])
			skills := append([]string{"普通攻擊"}, properties[4:]...)
			role := NewRole(properties[0], hp, mp, str, skills)
			troop.AddRole(role)
			continue
		}
		// TODO: 指令轉換成 行動準則
		fmt.Printf("run command %v\n", s)
	}
	// 開始戰鬥
	battle.Start()
}
