package command

import (
	"3-2/internal/item"
)

type MoveTankForward struct {
	tank item.Tank
}

func NewMoveTankForward(tank item.Tank) *MoveTankForward {
	return &MoveTankForward{tank: tank}
}

func (cmd MoveTankForward) Execute() {
	cmd.tank.MoveForward()
}

func (cmd MoveTankForward) Undo() {
	cmd.tank.BackForward()
}
