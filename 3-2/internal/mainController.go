package internal

type MainController struct {
	tank    Tank
	telecom Telecom
}

func NewMainController(tank Tank, telecom Telecom) *MainController {
	return &MainController{tank, telecom}
}

func (c *MainController) Press(key string) {
	switch key {
	case "f":
		c.tank.MoveForward()
	case "b":
		c.tank.BackForward()
	case "d":
		c.telecom.Connect()
	case "c":
		c.telecom.Disconnect()
	}
}
