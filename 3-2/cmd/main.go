package main

import (
	internal2 "3-2/internal"
	"3-2/internal/item"
	"log"
)

func main() {
	log.SetFlags(0)

	controller := internal2.NewMainController(item.Tank{}, item.Telecom{})
	controller.Start()
}
