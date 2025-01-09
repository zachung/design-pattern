package internal

type Troop struct {
	I      int
	roles  []*Role
	battle *Battle
}

func (t *Troop) AddRole(role *Role) {
	role.troop = t
	t.roles = append(t.roles, role)
}

func (t *Troop) AliveRoles() []*Role {
	roles := make([]*Role, 0)
	for _, role := range t.roles {
		if !role.IsDead() {
			roles = append(roles, role)
		}
	}
	return roles
}

func (t *Troop) IsAnnihilated() bool {
	return len(t.AliveRoles()) <= 0
}
