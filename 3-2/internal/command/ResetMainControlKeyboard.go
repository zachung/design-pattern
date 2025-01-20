package command

import (
	"3-2/internal/contract"
)

type ResetMainControlKeyboard struct {
	controller contract.Controller
	keyboard   map[string][]int
}

func NewResetMainControlKeyboard(c contract.Controller, keyboard map[string][]int) *ResetMainControlKeyboard {
	return &ResetMainControlKeyboard{controller: c, keyboard: keyboard}
}

func (cmd ResetMainControlKeyboard) Execute() {
	cmd.controller.Reset()
	cmd.controller.DisplayBinding()
}

func (cmd ResetMainControlKeyboard) Undo() {
	cmd.controller.SetKeyboard(cmd.keyboard)
	cmd.controller.DisplayBinding()
}
