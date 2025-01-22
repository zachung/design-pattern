package internal

import (
	"3-2/internal/item"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"sync"
	"testing"
)

func TestMainController(t *testing.T) {
	logWriter := new(strings.Builder)
	log.SetOutput(logWriter)
	log.SetFlags(0)
	input := make(chan string)

	wg := new(sync.WaitGroup)
	controller := NewMainController(item.Tank{}, item.Telecom{})
	wg.Add(1)
	go func() {
		controller.Start(input)
		wg.Done()
	}()

	inputs := []string{
		"1", "n", "f", "0", "f", "f", "f", "2", "2", "3", "3", "1", "y", "a", "0 2",
		"a", "2", "3", "1", "n", "r", "4", "r", "2", "3",
	}
	for _, s := range inputs {
		input <- s
	}
	close(input)
	wg.Wait()

	expectedOutput := `(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
1
設置巨集指令 (y/n)：
n
Key:
f
要將哪一道指令設置到快捷鍵 f 上:
(0) MoveTankForward 
(1) MoveTankBackward 
(2) ConnectTelecom 
(3) DisconnectTelecom 
(4) ResetMainControlKeyboard 
0
f: MoveTankForward
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
f
The tank has moved forward.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
f
The tank has moved forward.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
f
The tank has moved forward.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
2
The tank has moved backward.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
2
The tank has moved backward.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
3
The tank has moved forward.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
3
The tank has moved forward.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
1
設置巨集指令 (y/n)：
y
Key:
a
要將哪些指令設置成快捷鍵 a 的巨集（輸入多個數字，以空白隔開）:
(0) MoveTankForward 
(1) MoveTankBackward 
(2) ConnectTelecom 
(3) DisconnectTelecom 
(4) ResetMainControlKeyboard 
0 2
a: MoveTankForward & ConnectTelecom
f: MoveTankForward
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
a
The tank has moved forward.
The telecom has been turned on.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
2
The telecom has been turned off.
The tank has moved backward.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
3
The tank has moved forward.
The telecom has been turned on.
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
1
設置巨集指令 (y/n)：
n
Key:
r
要將哪一道指令設置到快捷鍵 r 上:
(0) MoveTankForward 
(1) MoveTankBackward 
(2) ConnectTelecom 
(3) DisconnectTelecom 
(4) ResetMainControlKeyboard 
4
a: MoveTankForward & ConnectTelecom
f: MoveTankForward
r: ResetMainControlKeyboard
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
r
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
2
a: MoveTankForward & ConnectTelecom
f: MoveTankForward
r: ResetMainControlKeyboard
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
3
(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵:
`
	expectedLines := strings.Split(expectedOutput, "\n")
	actualLines := strings.Split(logWriter.String(), "\n")
	for i, line := range expectedLines {
		if i >= len(actualLines) {
			assert.Failf(t, "missing at actual", "%s missed", line)
			break
		}
		assert.Equal(t, line, actualLines[i])
	}
}
