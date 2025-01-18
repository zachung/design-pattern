package command

import (
	"3-2/internal/item"
)

type MoveTankBackward struct {
	tank item.Tank
}

func NewMoveTankBackward(tank item.Tank) *MoveTankBackward {
	return &MoveTankBackward{tank: tank}
}

func (cmd MoveTankBackward) Execute() {
	cmd.tank.BackForward()
}

func (cmd MoveTankBackward) Undo() {
	cmd.tank.MoveForward()
}
