package internal

import (
	"3B/internal/contract"
	"fmt"
)

type Battle struct {
	troops []contract.Troop
	hero   contract.Role
}

func (b *Battle) Start(hero contract.Role) {
	fmt.Println("Starting Battle")
	b.hero = hero
	for {
		for _, role := range b.troops[0].AliveRoles() {
			role.Action(b.troops[1])
			if b.IsEnd() {
				return
			}
		}
		for _, role := range b.troops[1].AliveRoles() {
			role.Action(b.troops[0])
			if b.IsEnd() {
				return
			}
		}
	}
}

func (b *Battle) IsEnd() bool {
	if b.hero.IsDead() {
		fmt.Println("你失敗了！")
		return true
	}
	if b.troops[1].IsAnnihilated() {
		fmt.Println("你獲勝了！")
		return true
	}
	return false
}
