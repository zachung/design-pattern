package internal

import "fmt"

type Battle struct {
	troops []*Troop
}

func (b *Battle) Start() {
	fmt.Println("Starting Battle")
	for {
		for _, role := range b.troops[0].roles {
			if !role.IsDead() {
				if b.IsEnd() {
					return
				}
				role.Action(b.troops[1])
			}
		}
		for _, role := range b.troops[1].roles {
			if !role.IsDead() {
				if b.IsEnd() {
					return
				}
				role.Action(b.troops[0])
			}
		}
		fmt.Println("========")
	}
}

func (b *Battle) IsEnd() bool {
	if b.troops[0].roles[0].IsDead() {
		fmt.Println("你失敗了！")
		return true
	}
	if b.troops[1].IsAnnihilated() {
		fmt.Println("你獲勝了！")
		return true
	}
	return false
}
