package main

import (
	"3B/internal"
)

func main() {
	var initConfig = []string{
		"#軍隊-1-開始",
		"英雄 300 500 100 火球 水球",
		"#軍隊-1-結束",
		"#軍隊-2-開始",
		"Slime1 200 60 49 火球",
		"Slime2 200 200 50 火球 水球",
		"#軍隊-2-結束",
		"1",
		"2",
		"1",
		"2",
		"1",
		"2",
		"1",
	}
	internal.Run(initConfig)
}
