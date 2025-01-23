package internal

import (
	"3-B/internal/contract"
	"log"
)

type Battle struct {
	troops []contract.Troop
	hero   contract.Role
}

func NewBattle() *Battle {
	return &Battle{
		troops: make([]contract.Troop, 0),
	}
}

func (b *Battle) AddTroop(troop contract.Troop) {
	b.troops = append(b.troops, troop)
}

func (b *Battle) Start(hero contract.Role) {
	b.hero = hero
	for {
		if b.round(0, 1) {
			return
		}
		if b.round(1, 0) {
			return
		}
	}
}

func (b *Battle) round(team1Index, team2Index int) bool {
	i := 0
	for {
		role := b.troopAction(b.troops[team1Index], i)
		if role == nil {
			return false
		}
		if !role.IsDead() {
			role.Action(b.troops[team2Index])
			if b.isEnd() {
				return true
			}
		}
		i++
	}
}

func (b *Battle) troopAction(team1 contract.Troop, cur int) (role contract.Role) {
	for {
		team1 := team1.(*TroopImpl)
		if cur >= len(team1.roles) {
			return nil
		}
		return team1.roles[cur]
	}
}

func (b *Battle) isEnd() bool {
	if b.hero.IsDead() {
		log.Println("你失敗了！")
		return true
	}
	if b.troops[1].IsAnnihilated() {
		log.Println("你獲勝了！")
		return true
	}
	return false
}
