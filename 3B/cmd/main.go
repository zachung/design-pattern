package main

import (
	"3B/internal"
)

func main() {
	var initConfig = []string{
		"#軍隊-1-開始",
		"英雄 500 10000 30 召喚",
		"#軍隊-1-結束",
		"#軍隊-2-開始",
		"Slime1 1000 0 99",
		"#軍隊-2-結束",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
	}
	internal.Run(initConfig)
}
