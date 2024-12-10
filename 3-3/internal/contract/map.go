package contract

const (
	X                    = 10
	Y                    = 10
	MonsterName          = "怪物"
	CharacterName        = "玩家"
	TreasureName         = "寶物"
	ObstacleName         = "障礙物"
	CharacterUpSymbol    = '↑'
	CharacterRightSymbol = '→'
	CharacterDownSymbol  = '↓'
	CharacterLeftSymbol  = '←'
	ObstacleSymbol       = '□'
	MonsterSymbol        = 'M'
	TreasureSymbol       = 'x'
)

var Directions = map[rune]rune{
	'a': CharacterLeftSymbol,
	'w': CharacterUpSymbol,
	's': CharacterDownSymbol,
	'd': CharacterRightSymbol,
}

type Map interface {
	Draw()
	PutInRandomLocation(o Object)
	MoveObject(role *Role, direction rune)
	DestroyObject(object *Object)
	GetObjects() [][]*Object
	GetRoles() []*Role
}

type Object interface {
	GetName() string
	Symbol() rune
	Location() *Location
	SetLocation(location *Location)
}

type Role interface {
	Object
	GetHealth() int
	GetMaxHealth() int
	SubHealth(damage int)
	AddHealth(hp int)
	SetState(state State)
	Action()
	DoAction()
	Move()
	Attack()
}

type State interface {
	GetStateName() string
	OnRoundStart(m *Map)
	OnRounding(m *Map)
	OnRoundEnd(m *Map)
	OnTakeDamage()
	OnAttack(m *Map)
	GetAllowDirections() map[rune]rune
}
