package skill

import (
	"3-B/internal/contract"
	"fmt"
	"log"
	"strings"
)

type Curse struct {
	name   string
	mpCost int
}

// 紀錄 誰(map[Role]) 被 哪些角色([]Role)詛咒
var curses map[contract.Role][]contract.Role

func init() {
	curses = make(map[contract.Role][]contract.Role)
}

func NewCurse() contract.Skill {
	return &Curse{
		name:   "詛咒",
		mpCost: 100,
	}
}

func (a *Curse) GetName() string {
	return a.name
}

func (a *Curse) CanCast(role contract.Role) bool {
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *Curse) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
	targetCount := 1
	targets := enemy.AliveRoles()
	if targetCount < len(targets) {
		targets = role.Actor().SelectTarget(targets, targetCount)
	}
	role.Property(contract.Mp).Sub(a.mpCost)

	var str []string
	for _, enemy := range targets {
		str = append(str, enemy.GetName())
	}
	log.Println(fmt.Sprintf("%s 對 %s 使用了 %s。", role.GetName(), strings.Join(str, ", "), a.GetName()))
	for _, enemy := range targets {
		curses[enemy] = append(curses[enemy], role)
		enemy.Property(contract.Hp).AddObserver(func(hp *int) {
			if *hp <= 0 {
				mp := enemy.Property(contract.Mp).Get()
				for _, curseFrom := range curses[enemy] {
					curseFrom.Property(contract.Hp).Add(mp)
				}
				// 清空已處理詛咒
				curses[enemy] = nil
			}
		})
	}
}
