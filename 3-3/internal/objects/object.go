package objects

import (
	"3-3/internal/contract"
	"3-3/internal/states"
	"fmt"
)

type GameObject struct {
	contract.Object
	location *contract.Location
	gameMap  *contract.Map
	name     string
}

func (o *GameObject) Location() *contract.Location {
	return o.location
}

func (o *GameObject) SetLocation(location *contract.Location) {
	o.location = location
}

func (o *GameObject) GetName() string {
	return o.name
}

type RoleStruct struct {
	GameObject
	contract.State
	health    int
	maxHealth int
}

func (r *RoleStruct) GetHealth() int {
	return r.health
}

func (r *RoleStruct) GetMaxHealth() int {
	return r.maxHealth
}

func (r *RoleStruct) Action() {
	r.State.OnRoundStart(r.gameMap)
	r.State.OnRounding(r.gameMap)
	r.State.OnRoundEnd(r.gameMap)
}

func (r *RoleStruct) DoAction() {
}

func (r *RoleStruct) Move() {
}

func (r *RoleStruct) SubHealth(damage int) {
	name := r.GetName()
	switch r.State.(type) {
	case *states.Invincible:
		fmt.Printf("%s是無敵的\n", name)
		// 無敵狀態
		return
	}
	r.health -= damage
	fmt.Printf("%s受到%d傷害\n", name, damage)
	r.OnTakeDamage()
	if r.health <= 0 {
		object := contract.Object(r)
		(*r.gameMap).DestroyObject(&object)
	} else {
		r.OnTakeDamage()
	}
}

func (r *RoleStruct) AddHealth(hp int) {
	r.health += hp
	if r.health >= r.maxHealth {
		r.health = r.maxHealth
	}
}

func (r *RoleStruct) SetState(state contract.State) {
	r.State = state
}
