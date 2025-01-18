package command

import (
	"3-2/internal/item"
)

type ConnectTelecom struct {
	telecom item.Telecom
}

func NewConnectTelecom(telecom item.Telecom) *ConnectTelecom {
	return &ConnectTelecom{telecom: telecom}
}

func (cmd ConnectTelecom) Execute() {
	cmd.telecom.Connect()
}

func (cmd ConnectTelecom) Undo() {
	cmd.telecom.Disconnect()
}
