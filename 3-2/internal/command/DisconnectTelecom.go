package command

import (
	"3-2/internal/item"
)

type DisconnectTelecom struct {
	telecom item.Telecom
}

func NewDisconnectTelecom(telecom item.Telecom) *DisconnectTelecom {
	return &DisconnectTelecom{telecom: telecom}
}

func (cmd DisconnectTelecom) Execute() {
	cmd.telecom.Disconnect()
}

func (cmd DisconnectTelecom) Undo() {
	cmd.telecom.Connect()
}
