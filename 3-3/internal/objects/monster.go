package objects

import (
	"3-3/internal/contract"
	"3-3/internal/states"
	"fmt"
)

type Monster struct {
	RoleStruct
}

func NewMonster(gameMap *contract.Map) *Monster {
	monster := &Monster{
		RoleStruct: RoleStruct{
			GameObject: GameObject{
				gameMap: gameMap,
				name:    contract.MonsterName,
			},
			health:    1,
			maxHealth: 1,
		},
	}
	role := contract.Role(monster)
	monster.State = *states.NewNormal(&role)

	return monster
}

func (m *Monster) Symbol() rune {
	return contract.MonsterSymbol
}

func (m *Monster) DoAction() {
	m.Attack()
}

func (m *Monster) getNearCharacter() *Character {
	// 檢查是否在附近
	mLocation := *m.Location()
	locations := []*contract.Location{
		mLocation.Copy().Up(),
		mLocation.Copy().Down(),
		mLocation.Copy().Left(),
		mLocation.Copy().Right(),
	}
	objects := (*m.gameMap).GetObjects()
	for _, location := range locations {
		object := objects[location.Y][location.X]
		if object == nil {
			continue
		}
		switch (*object).(type) {
		case *Character:
			return (*object).(*Character)
		}
	}

	return nil
}

func (m *Monster) Attack() {
	character := m.getNearCharacter()
	if character == nil {
		m.Move()
		return
	}
	fmt.Printf("%s攻擊%s\n", m.GetName(), character.GetName())
	character.SubHealth(50)
}

func (m *Monster) Move() {
	for _, direction := range m.State.GetAllowDirections() {
		// use map rand
		role := contract.Role(m)
		(*m.gameMap).MoveObject(&role, direction)
		return
	}
}
