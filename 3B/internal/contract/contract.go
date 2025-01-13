package contract

type Troop interface {
	GetI() int
	NewRole(data string) Role
	AddRole(role Role)
	AliveRoles() []Role
	IsAnnihilated() bool
}

type Role interface {
	SetActor(actor Actor)
	IsDead() bool
	GetName() string
	Action(targetTroop Troop)
	SelectSkill(selected int) Skill
	SelectTarget(enemies []Role, targetCount int, selected []int) (targets []Role)
	Actor() Actor
	SetState(state State)
	SetObserver(event Event, observer func())
	MakeDamage(damage int) int
	Property(key PropertyKey) Property
	GetState() State
}

type Skill interface {
	GetName() string
	CanCast(role Role) bool
	Cast(role Role, ally Troop, enemy Troop)
}

type Actor interface {
	SelectSkill(skillCount int) Skill
	SelectTarget(enemies []Role, targetCount int) []Role
	CastSkill(s Skill, ally Troop, enemy Troop)
}

type State interface {
	GetName() string
	CanAction() bool
	BeforeAction()
	AfterAction()
	IsFinished() bool
	MakeDamage(damage int) int
}

type Event int

const (
	OnDead Event = iota
)

type PropertyKey int

const (
	Hp PropertyKey = iota
	Mp
	Str
)

type Property interface {
	Get() int
	Sub(v int)
	Add(v int)
	AddObserver(observer func(*int))
}
