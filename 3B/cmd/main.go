package main

import (
	"3B/cmd/client"
	"fmt"
	"log"
	"strings"
)

func main() {
	input := `#軍隊-1-開始
英雄 500 500 40
#軍隊-1-結束
#軍隊-2-開始
Slime1 100 100 30 自我治療
#軍隊-2-結束
0
0
0
0
0
0
0
0
0
0`

	logWriter := new(strings.Builder)
	log.SetOutput(logWriter)
	log.SetFlags(0)

	game := client.NewGame(strings.Split(input, "\n"))
	game.Start()

	logs := logWriter.String()
	fmt.Println(logs)
}
