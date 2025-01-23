package internal

import (
	"3-B/internal/contract"
	"strconv"
	"strings"
)

type TroopImpl struct {
	I     int
	roles []contract.Role
}

func (t *TroopImpl) GetI() int {
	return t.I
}

func (t *TroopImpl) NewRole(data string) contract.Role {
	properties := strings.Split(data, " ")
	name := properties[0]
	hp, _ := strconv.Atoi(properties[1])
	mp, _ := strconv.Atoi(properties[2])
	str, _ := strconv.Atoi(properties[3])
	skills := append([]string{"普通攻擊"}, properties[4:]...)
	role := NewRole(t, name, hp, mp, str, skills)

	return role
}

func (t *TroopImpl) AddRole(role contract.Role) {
	t.roles = append(t.roles, role)
}

func (t *TroopImpl) AliveRoles() []contract.Role {
	roles := make([]contract.Role, 0)
	for _, role := range t.roles {
		if !role.IsDead() {
			roles = append(roles, role)
		}
	}
	return roles
}

func (t *TroopImpl) IsAnnihilated() bool {
	return len(t.AliveRoles()) <= 0
}
