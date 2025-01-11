package internal

import "3B/internal/contract"

type TroopImpl struct {
	I      int
	roles  []contract.Role
	battle *Battle
}

func (t *TroopImpl) GetI() int {
	return t.I
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
