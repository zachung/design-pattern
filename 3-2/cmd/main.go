package main

import (
	internal2 "3-2/internal"
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)

	controller := internal2.NewMainController(internal2.Tank{}, internal2.Telecom{})
	for {
		log.Println("按下按鍵: ")
		controller.Press(readTypeIn())
	}
}

func readTypeIn() string {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		return text
	}
}
