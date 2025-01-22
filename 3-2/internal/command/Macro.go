package command

import (
	"3-2/internal/contract"
)

type Macro struct {
	commands []contract.Command
}

func NewMacro(commands []contract.Command) *Macro {
	return &Macro{commands: commands}
}

func (cmd Macro) Execute() {
	for _, command := range cmd.commands {
		command.Execute()
	}
}

func (cmd Macro) Undo() {
	for i := len(cmd.commands) - 1; i >= 0; i-- {
		cmd.commands[i].Undo()
	}
}
