package main

import (
	"3B/internal"
)

func main() {
	var initConfig = []string{
		"#軍隊-1-開始",
		"英雄 500 10000 30 鼓舞",
		"Servant1 1000 0 0",
		"Servant2 1000 0 0",
		"Servant3 1000 0 0",
		"Servant4 1000 0 0",
		"Servant5 1000 0 0",
		"#軍隊-1-結束",
		"#軍隊-2-開始",
		"Slime1 500 0 0",
		"#軍隊-2-結束",
		"1",
		"0, 1, 2",
		"1",
		"2, 3, 4",
		"0",
	}
	internal.Run(initConfig)
}
