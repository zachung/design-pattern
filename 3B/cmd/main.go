package main

import (
	"3B/internal"
)

func main() {
	var initConfig = []string{
		"#軍隊-1-開始",
		"英雄 500 500 40",
		"#軍隊-1-結束",
		"#軍隊-2-開始",
		"Slime1 100 100 30 自我治療",
		"#軍隊-2-結束",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
	}
	internal.Run(initConfig)
}
