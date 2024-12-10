package internal

import (
	"3-3/internal/contract"
	"3-3/internal/objects"
	"fmt"
	"math/rand/v2"
)

type Map struct {
	objects   [][]*contract.Object
	monsters  []*objects.Monster
	treasures []*objects.Treasure
	character *objects.Character
}

func NewMap() *Map {
	// init map tiles
	objs := make([][]*contract.Object, contract.Y)
	for i := range objs {
		objs[i] = make([]*contract.Object, contract.X)
	}
	m := Map{objects: objs}
	im := contract.Map(&m)

	// init obstacles
	for i := 0; i < 10; i++ {
		obstacle := objects.NewObstacle(&im)
		m.PutInRandomLocation(contract.Object(obstacle))
	}

	// init treasures
	for i := 0; i < 10; i++ {
		treasure := objects.NewTreasure(&im)
		m.PutInRandomLocation(contract.Object(treasure))
	}

	// init monsters
	m.monsters = make([]*objects.Monster, 1)
	for i := 0; i < len(m.monsters); i++ {
		monster := objects.NewMonster(&im)
		m.PutInRandomLocation(contract.Object(monster))
		m.monsters[i] = monster
	}

	// init character
	m.character = objects.NewCharacter(&im)
	m.PutInRandomLocation(contract.Object(m.character))

	return &m
}

func (m *Map) PutInRandomLocation(o contract.Object) {
	for {
		preLocation := o.Location()
		location := newRandLocation()
		if m.objects[location.Y][location.X] == nil {
			o.SetLocation(location)
			m.objects[location.Y][location.X] = &o
			if preLocation != nil {
				m.objects[preLocation.Y][preLocation.X] = nil
			}
			break
		}
	}
}

func (m *Map) Start() {
	for {
		m.Draw()
		if m.character.GetHealth() <= 0 {
			fmt.Println("Game Over")
			return
		}
		if len(m.monsters) == 0 {
			fmt.Println("You win!")
			return
		}
		m.character.Action()
		for _, monster := range m.monsters {
			monster.Action()
		}
		m.spawn()
	}
}

func (m *Map) Draw() {
	for i := range m.objects {
		for j := range m.objects[i] {
			if m.objects[i][j] == nil {
				fmt.Printf("%s", " ")
			} else {
				fmt.Printf("%s", string((*m.objects[i][j]).Symbol()))
			}
		}
		fmt.Println()
	}
	fmt.Printf("生命: %d, 狀態: %s\n", m.character.GetHealth(), m.character.GetStateName())
}

func (m *Map) MoveObject(role *contract.Role, direction rune) {
	r := *role
	departL := r.Location()
	destL := departL.Copy()
	switch direction {
	case contract.CharacterRightSymbol:
		destL.Right()
	case contract.CharacterLeftSymbol:
		destL.Left()
	case contract.CharacterDownSymbol:
		destL.Down()
	case contract.CharacterUpSymbol:
		destL.Up()
	}
	if departL == destL {
		// 沒動
		return
	}
	current := contract.Object(r)
	target := m.objects[destL.Y][destL.X]
	if target == nil {
		m.objects[departL.Y][departL.X] = nil
		m.objects[destL.Y][destL.X] = &current
		r.SetLocation(destL)
		fmt.Printf("%s往%s移動\n", r.GetName(), string(direction))
		return
	}
	// touched
	m.touch(role, target)
}

func (m *Map) DestroyObject(object *contract.Object) {
	location := (*object).Location()
	o := *m.objects[location.Y][location.X]
	switch o.(type) {
	case *objects.Monster:
		for i, monster := range m.monsters {
			if monster == o {
				m.monsters = append(m.monsters[:i], m.monsters[i+1:]...)
				break
			}
		}
	case *objects.Treasure:
		for i, treasure := range m.treasures {
			if treasure == o {
				m.treasures = append(m.treasures[:i], m.treasures[i+1:]...)
			}
		}
	}
	m.objects[location.Y][location.X] = nil
}

func (m *Map) touch(current *contract.Role, target *contract.Object) {
	switch (*target).(type) {
	case *objects.Treasure:
		treasure := (*target).(*objects.Treasure)
		(*current).SetState(*treasure.GetState(current))
		fmt.Printf("%s獲得%s\n", (*current).GetName(), treasure.GetTreasureName())
		m.DestroyObject(target)
	}
}

func newRandLocation() *contract.Location {
	return &contract.Location{X: rand.UintN(contract.X), Y: rand.UintN(contract.Y)}
}

func (m *Map) spawn() {
	im := contract.Map(m)
	// spawn treasure
	tN := rand.UintN(100)
	if tN < 20 {
		treasure := objects.NewTreasure(&im)
		m.PutInRandomLocation(contract.Object(treasure))
		fmt.Println("一個新寶物生成")
	}

	// spawn monster
	mN := rand.UintN(100)
	if mN < 20 {
		monster := objects.NewMonster(&im)
		m.PutInRandomLocation(contract.Object(monster))
		m.monsters = append(m.monsters, monster)
		fmt.Println("一個新怪物生成")
	}
}

func (m *Map) GetObjects() [][]*contract.Object {
	return m.objects
}

func (m *Map) GetRoles() []*contract.Role {
	roles := make([]*contract.Role, 0)
	for _, monster := range m.monsters {
		role := contract.Role(monster)
		roles = append(roles, &role)
	}
	role := contract.Role(m.character)
	roles = append(roles, &role)

	return roles
}
