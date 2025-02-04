package internal

import (
	"3-2/internal/command"
	"3-2/internal/contract"
	"3-2/internal/item"
	"log"
	"strconv"
	"strings"
	"testing"
)

var availableCmds = []string{
	0: "MoveTankForward",
	1: "MoveTankBackward",
	2: "ConnectTelecom",
	3: "DisconnectTelecom",
	4: "ResetMainControlKeyboard",
}

type MainController struct {
	reader       <-chan string
	tank         item.Tank
	telecom      item.Telecom
	keyboard     map[string][]int
	history      []contract.Command
	historyIndex int
}

func NewMainController(tank item.Tank, telecom item.Telecom) *MainController {
	return &MainController{
		tank:         tank,
		telecom:      telecom,
		keyboard:     make(map[string][]int),
		historyIndex: 0,
	}
}

func (c *MainController) Start(reader <-chan string) {
	c.reader = reader
	log.Println("(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:")
	for in := range c.reader {
		if testing.Testing() {
			log.Println(in)
		}
		switch in {
		case "1":
			c.bind()
		case "2":
			c.undo()
		case "3":
			c.redo()
		default:
			c.press(in)
		}
		log.Println("(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:")
	}
}

func (c *MainController) press(key string) {
	cmds, ok := c.keyboard[key]
	if !ok {
		return
	}
	var cmd contract.Command
	if len(cmds) == 1 {
		i := cmds[0]
		cmd = c.getCommand(i)
	} else {
		var commands []contract.Command
		for _, i := range cmds {
			commands = append(commands, c.getCommand(i))
		}
		cmd = command.NewMacro(commands)
	}
	cmd.Execute()
	c.history = append(c.history[:c.historyIndex], cmd)
	c.historyIndex = len(c.history)
}

func (c *MainController) bind() {
	var isMacro bool
	for {
		log.Printf("設置巨集指令 (y/n)：")
		in := c.readTypeIn()
		if in == "y" {
			isMacro = true
			break
		} else if in == "n" {
			isMacro = false
			break
		}
	}
	log.Printf("Key:")
	key := c.readTypeIn()
	if isMacro {
		log.Printf("要將哪些指令設置成快捷鍵 %s 的巨集（輸入多個數字，以空白隔開）:\n", key)
	} else {
		log.Printf("要將哪一道指令設置到快捷鍵 %s 上:\n", key)
	}
	for i, cmd := range availableCmds {
		log.Printf("(%d) %s ", i, cmd)
	}
	var cmdIndexes []int
	s := c.readTypeIn()
	cmds := strings.Split(s, " ")
	for _, cmd := range cmds {
		i, _ := strconv.Atoi(cmd)
		cmdIndexes = append(cmdIndexes, i)
	}
	c.keyboard[key] = cmdIndexes
	c.DisplayBinding()
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
	c.keyboard = make(map[string][]int)
}

func (c *MainController) SetKeyboard(keyboard map[string][]int) {
	c.keyboard = keyboard
}

func (c *MainController) DisplayBinding() {
	// a-z
	for i := range 26 {
		s := string(rune(i + 97))
		if c.keyboard[s] != nil {
			ints := c.keyboard[s]
			listenFuncs := availableCmds[ints[0]]
			for i := 1; i < len(ints); i++ {
				listenFuncs += " & " + availableCmds[ints[i]]
			}
			log.Printf("%s: %s", s, listenFuncs)
		}
	}
}

func (c *MainController) getCommand(i int) (cmd contract.Command) {
	pairCmd := availableCmds[i]
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
	}
	return
}

func (c *MainController) readTypeIn() string {
	s := <-c.reader
	if testing.Testing() {
		log.Println(s)
	}
	return s
}
