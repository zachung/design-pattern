package main

import (
	"3B/internal"
)

func main() {
	var initConfig = []string{
		"#軍隊-1-開始",
		"英雄 1000 10000 0 一拳攻擊 下毒 石化 鼓舞",
		"#軍隊-1-結束",
		"#軍隊-2-開始",
		"Slime1 601 0 0",
		"Slime2 241 0 0",
		"Slime3 101 999 0 一拳攻擊 一拳攻擊 鼓舞",
		"#軍隊-2-結束",
		"1",
		"0",
		"2",
		"1",
		"1",
		"1",
		"1",
		"0",
		"1",
		"0",
		"3",
		"0",
		"1",
		"0",
		"2",
		"1",
	}
	internal.Run(initConfig)
}
