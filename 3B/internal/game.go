package internal

import (
	"3B/internal/contract"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	battle Battle
	hero   contract.Role
}

func NewGame(initConfig []string) *Game {
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
			if troop == nil {
				panic("軍隊還沒建立")
			}
			// 添加角色
			role := troop.NewRole(s)
			if strings.Contains(role.GetName(), "英雄") {
				hero = role
				role.SetActor(&HeroAction{role: role, controller: controller})
			}
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
	return &Game{
		battle: battle,
		hero:   hero,
	}
}

func (g *Game) Start() {
	g.battle.Start(g.hero)
}
