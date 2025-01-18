package internal

import (
	"3-2/internal/command"
	"3-2/internal/contract"
	"3-2/internal/item"
	"3-2/internal/utils"
	"log"
	"strconv"
)

var availableCmds = []string{
	0: "MoveTankForward",
	1: "MoveTankBackward",
	2: "ConnectTelecom",
	3: "DisconnectTelecom",
	4: "ResetMainControlKeyboard",
}

type MainController struct {
	tank         item.Tank
	telecom      item.Telecom
	keyboard     map[string]int
	history      []contract.Command
	historyIndex int
}

func NewMainController(tank item.Tank, telecom item.Telecom) *MainController {
	return &MainController{
		tank:         tank,
		telecom:      telecom,
		keyboard:     make(map[string]int),
		historyIndex: 0,
	}
}

func (c *MainController) Start() {
	for {
		log.Println("(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:")
		in := utils.ReadTypeIn()
		switch in {
		case "1":
			log.Printf("Key:")
			key := utils.ReadTypeIn()
			c.bind(key)
			c.DisplayBinding()
		case "2":
			c.undo()
		case "3":
			c.redo()
		default:
			c.press(in)
		}
	}
}

func (c *MainController) press(key string) {
	i, ok := c.keyboard[key]
	if !ok {
		return
	}
	pairCmd := availableCmds[i]
	var cmd contract.Command
	switch pairCmd {
	case "MoveTankForward":
		cmd = command.NewMoveTankForward(c.tank)
	case "MoveTankBackward":
		cmd = command.NewMoveTankBackward(c.tank)
	case "ConnectTelecom":
		cmd = command.NewConnectTelecom(c.telecom)
	case "DisconnectTelecom":
		cmd = command.NewDisconnectTelecom(c.telecom)
	case "ResetMainControlKeyboard":
		cmd = command.NewResetMainControlKeyboard(c, c.keyboard)
	default:
		return
	}
	cmd.Execute()
	c.history = append(c.history[:c.historyIndex], cmd)
	c.historyIndex = len(c.history)
}

func (c *MainController) bind(key string) {
	log.Printf("要將哪一道指令設置到快捷鍵 %s 上:\n", key)
	for i, cmd := range availableCmds {
		log.Printf("(%d) %s ", i, cmd)
	}
	var cmdIndex int
	for {
		i, _ := strconv.Atoi(utils.ReadTypeIn())
		if i >= 0 && i < len(availableCmds) {
			cmdIndex = i
			break
		}
	}
	c.keyboard[key] = cmdIndex
}

func (c *MainController) undo() {
	if c.historyIndex <= 0 {
		return
	}
	c.history[c.historyIndex-1].Undo()
	c.historyIndex--
}

func (c *MainController) redo() {
	if c.historyIndex >= len(c.history) {
		return
	}
	c.history[c.historyIndex].Execute()
	c.historyIndex++
}

func (c *MainController) Reset() {
	c.keyboard = make(map[string]int)
}

func (c *MainController) SetKeyboard(keyboard map[string]int) {
	c.keyboard = keyboard
}

func (c *MainController) DisplayBinding() {
	// display key mapped
	for s, i := range c.keyboard {
		log.Printf("%s: %s", s, availableCmds[i])
	}
}
