package objects

import (
	"3-3/internal/contract"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func (c *Character) DoAction() {
	fmt.Println("請輸入指令: 1.往一個方向移動一格, 2.攻擊")
	action := rune(readTypeIn([]byte{'1', '2'}))
	switch action {
	case '1':
		c.Move()
	case '2':
		c.State.OnAttack(c.gameMap)
	}
}

func (c *Character) Move() {
	direction := c.GetDirection(c.State.GetAllowDirections())
	d := contract.Directions[direction]
	c.SetDirection(d)
	role := contract.Role(c)
	(*c.gameMap).MoveObject(&role, d)
}

func (c *Character) GetDirection(allow map[rune]rune) rune {
	var keys []byte
	for key := range allow {
		keys = append(keys, byte(key))
	}
	var s string
	for _, b := range keys {
		s += string(b) + ","
	}
	fmt.Printf("請輸入方向:(%s)\n", s[:len(s)-1])
	return rune(readTypeIn(keys))
}

func readTypeIn(bytes []byte) byte {
	for {
		f := bufio.NewReader(os.Stdin)
		b, err := f.ReadByte()
		if err != nil && !errors.Is(err, io.EOF) {
			continue
		}
		for _, v := range bytes {
			if b == v {
				return v
			}
		}
	}
}

func (c *Character) Attack() {
	direction := c.Symbol()
	location := c.Location().Copy()
	var moveFunc func() *contract.Location
	switch direction {
	case contract.CharacterUpSymbol:
		moveFunc = location.Up
	case contract.CharacterDownSymbol:
		moveFunc = location.Down
	case contract.CharacterLeftSymbol:
		moveFunc = location.Left
	case contract.CharacterRightSymbol:
		moveFunc = location.Right
	}
	m := *c.gameMap
	objects := m.GetObjects()
	curL := location.Copy()
	for {
		newL := moveFunc()
		if *newL == *curL {
			// 位置沒變
			return
		}
		if !c.attackTarget(objects[location.Y][location.X]) {
			return
		}
		curL = newL.Copy()
	}
}

func (c *Character) attackTarget(target *contract.Object) bool {
	if target == nil {
		return true
	}
	switch (*target).(type) {
	case *Obstacle:
		return false
	case *Monster:
		// attack target
		monster := (*target).(*Monster)
		fmt.Printf("%s攻擊%s\n", c.GetName(), monster.GetName())
		monster.SubHealth(1)
	}
	return true
}
