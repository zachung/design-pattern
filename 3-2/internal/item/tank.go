package item

import (
	"log"
)

type Tank struct {
}

func (t Tank) MoveForward() {
	log.Println("The tank has moved forward.")
}

func (t Tank) BackForward() {
	log.Println("The tank has moved backward.")
}
