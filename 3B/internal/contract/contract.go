package contract

type Troop interface {
	GetI() int
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
	SubHp(damage int)
	AddHp(hp int)
	GetStr() int
	GetMp() int
	SubMp(mp int)
	SetState(state State)
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
	AfterAction()
	IsFinished() bool
}
