package internal

import "fmt"

type Tank struct {
}

func (t Tank) MoveForward() {
	fmt.Println("The tank has moved forward.")
}

func (t Tank) BackForward() {
	fmt.Println("The tank has moved backward.")
}
