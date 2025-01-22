package main

import (
	internal2 "3-2/internal"
	"3-2/internal/item"
	"bufio"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	input := make(chan string)
	controller := internal2.NewMainController(item.Tank{}, item.Telecom{})
	go func() {
		defer close(input)
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			input <- scan.Text()
		}
	}()
	controller.Start(input)
}
