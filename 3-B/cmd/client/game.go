package client

import (
	"3-B/internal"
	"3-B/internal/actor"
	"3-B/internal/contract"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	battle *internal.Battle
	hero   contract.Role
}

func NewGame(input []string) *Game {
	var hero contract.Role
	var troop contract.Troop
	var battle = internal.NewBattle()
	var cmdStrs []string
	controller := actor.NewController()
	var countTroops int
	for _, s := range input {
		r := regexp.MustCompile(`#軍隊-(\d)-(.*)`)
		match := r.FindStringSubmatch(s)
		if len(match) > 1 {
			i, _ := strconv.Atoi(match[1])
			// 切換軍隊
			if match[2] == "開始" {
				troop = &internal.TroopImpl{I: i}
			}
			if match[2] == "結束" {
				battle.AddTroop(troop)
				countTroops++
			}
			continue
		}
		if countTroops < 2 {
			if troop == nil {
				panic("軍隊還沒建立")
			}
			// 添加角色
			role := troop.NewRole(s)
			if strings.Contains(role.GetName(), "英雄") {
				hero = role
				role.SetActor(actor.NewHeroAction(role, controller))
			} else {
				role.SetActor(actor.NewAI(role))
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
