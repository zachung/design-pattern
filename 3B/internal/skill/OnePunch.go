package skill

import (
	"3B/internal/contract"
	"3B/internal/state"
	"fmt"
	"log"
	"strings"
)

type OnePunch struct {
	name   string
	mpCost int
	damage int
}

func NewOnePunch() contract.Skill {
	return &OnePunch{
		name:   "一拳攻擊",
		mpCost: 180,
	}
}

func (a *OnePunch) GetName() string {
	return a.name
}

func (a *OnePunch) CanCast(role contract.Role) bool {
	return role.Property(contract.Mp).Get() >= a.mpCost
}

func (a *OnePunch) Cast(role contract.Role, ally contract.Troop, enemy contract.Troop) {
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
		a.applySkill(role, enemy)
	}
}

func (a *OnePunch) applySkill(role contract.Role, enemy contract.Role) {
	if enemy.Property(contract.Hp).Get() >= 500 {
		damage := role.MakeDamage(300)
		log.Println(fmt.Sprintf("%s 對 %s 造成 %d 點傷害。", role.GetName(), enemy.GetName(), damage))
		enemy.Property(contract.Hp).Sub(damage)
		return
	}
	stateName := enemy.GetState().GetName()
	if stateName == "中毒" || stateName == "石化" {
		damage := role.MakeDamage(80)
		for i := 0; i < 3; i++ {
			log.Println(fmt.Sprintf("%s 對 %s 造成 %d 點傷害。", role.GetName(), enemy.GetName(), damage))
			enemy.Property(contract.Hp).Sub(damage)
			if enemy.IsDead() {
				return
			}
		}
		return
	}
	if stateName == "受到鼓舞" {
		damage := role.MakeDamage(100)
		log.Println(fmt.Sprintf("%s 對 %s 造成 %d 點傷害。", role.GetName(), enemy.GetName(), damage))
		enemy.Property(contract.Hp).Sub(damage)
		enemy.SetState(state.NewNormalState())
		return
	}
	if stateName == "正常" {
		damage := role.MakeDamage(100)
		log.Println(fmt.Sprintf("%s 對 %s 造成 %d 點傷害。", role.GetName(), enemy.GetName(), damage))
		enemy.Property(contract.Hp).Sub(damage)
		return
	}
}
