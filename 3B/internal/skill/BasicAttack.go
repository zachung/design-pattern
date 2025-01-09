package skill

type Area int

type Targets int

const (
	Enemy Targets = iota
	Ally
)

type Skill struct {
	Name   string
	MpCost int
	Area   []int
	Damage int
}

var BasicAttack = Skill{Name: "普通攻擊", MpCost: 0, Area: []int{1, 0}}
var Waterball = Skill{Name: "水球", MpCost: 50, Area: []int{1, 0}, Damage: 120}
var Fireball = Skill{Name: "火球", MpCost: 50, Area: []int{-1, 0}, Damage: 50}
