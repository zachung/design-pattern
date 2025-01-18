package contract

type Controller interface {
	Reset()
	SetKeyboard(keyboard map[string]int)
	DisplayBinding()
}

type Command interface {
	Execute()
	Undo()
}
