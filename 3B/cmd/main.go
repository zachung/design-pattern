package main

import (
	"3B/internal"
)

func main() {
	var initConfig = []string{
		"#軍隊-1-開始",
		"英雄 300 10000 100 詛咒",
		"Ally 600 100 30 詛咒 詛咒",
		"#軍隊-1-結束",
		"#軍隊-2-開始",
		"Slime1 200 999 50",
		"Slime2 200 999 100",
		"#軍隊-2-結束",
		"1",
		"1",
		"0",
		"0",
		"0",
		"1",
		"0",
		"1",
		"0",
	}
	internal.Run(initConfig)
}
