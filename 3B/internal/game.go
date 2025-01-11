package internal

import (
	"3B/internal/contract"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	battle Battle
}

func Run(initConfig []string) {
	var hero contract.Role
	var troop contract.Troop
	var battle = Battle{troops: make([]contract.Troop, 0)}
	var cmdStrs []string
	controller := NewController()
	for _, s := range initConfig {
		r := regexp.MustCompile(`#軍隊-(\d)-(.*)`)
		match := r.FindStringSubmatch(s)
		if len(match) > 1 {
			i, _ := strconv.Atoi(match[1])
			// 切換軍隊
			if match[2] == "開始" {
				troop = &TroopImpl{I: i}
			}
			if match[2] == "結束" {
				battle.troops = append(battle.troops, troop)
			}
			continue
		}
		if len(battle.troops) < 2 {
			// 添加角色
			properties := strings.Split(s, " ")
			name := properties[0]
			hp, _ := strconv.Atoi(properties[1])
			mp, _ := strconv.Atoi(properties[2])
			str, _ := strconv.Atoi(properties[3])
			skills := append([]string{"普通攻擊"}, properties[4:]...)
			role := NewRole(troop, name, hp, mp, str, skills)
			var actor contract.Actor
			if name == "英雄" {
				actor = &Round{role: role, controller: controller}
				hero = role
			} else {
				actor = &AI{role: role}
			}
			role.SetActor(actor)
			troop.AddRole(role)
			continue
		}
		// 行動準則
		cmdStrs = append(cmdStrs, s)
	}
	go func() {
		for _, s := range cmdStrs {
			cmds := strings.Split(s, ", ")
			ints := make([]int, len(cmds))
			for i, cmd := range cmds {
				ints[i], _ = strconv.Atoi(cmd)
			}
			controller.AddCommand(ints)
		}
	}()
	// 開始戰鬥
	battle.Start(hero)
}
