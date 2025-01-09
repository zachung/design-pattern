package skill

type Area int

const (
	Enemy Area = iota
	Ally
	AllEnemy
	AllAlly
	All
)

type Targets int

type Skill struct {
	MpCost int
	Area
	Targets
	Damage int
}

var BasicAttack = Skill{MpCost: 0, Area: Enemy, Targets: 1}
var Waterball = Skill{MpCost: 50, Area: Enemy, Targets: 1, Damage: 120}
var Fireball = Skill{MpCost: 50, Area: AllEnemy, Damage: 50}
