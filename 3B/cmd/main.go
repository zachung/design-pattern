package main

import (
	"3B/internal"
)

func main() {
	var initConfig = []string{
		"#軍隊-1-開始",
		"英雄 1000 500 0 下毒",
		"#軍隊-1-結束",
		"#軍隊-2-開始",
		"Slime1 120 90 50",
		"Slime2 120 90 50",
		"Slime3 120 9000 50",
		"#軍隊-2-結束",
		"1",
		"0",
		"1",
		"1",
		"1",
		"2",
		"1",
		"0",
		"1",
		"1",
		"1",
		"0",
	}
	internal.Run(initConfig)
}
