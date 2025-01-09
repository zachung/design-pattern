package internal

import (
	"strconv"
	"strings"
)

type Controller struct {
	commands [][]int
}

func NewController() *Controller {
	return &Controller{commands: make([][]int, 0)}
}

func (c *Controller) AddCommand(cmd string) {
	cmds := strings.Split(cmd, ", ")
	ints := make([]int, len(cmds))
	for i, s := range cmds {
		ints[i], _ = strconv.Atoi(s)
	}
	c.commands = append(c.commands, ints)
}

func (c *Controller) PullCommand() []int {
	var command []int
	command, c.commands = c.commands[0], c.commands[1:]

	return command
}
