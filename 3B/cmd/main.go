package main

import (
	"3B/internal"
)

func main() {
	var initConfig = []string{
		"#軍隊-1-開始",
		"英雄 400 99999 30 石化",
		"#軍隊-1-結束",
		"#軍隊-2-開始",
		"攻擊力超強的BOSS 270 9999 399 石化",
		"#軍隊-2-結束",
		"1",
		"0",
		"0",
		"0",
		"1",
		"0",
		"0",
		"1",
		"0",
		"0",
		"0",
	}
	internal.Run(initConfig)
}
