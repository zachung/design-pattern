package item

import (
	"log"
)

type Telecom struct{}

func (t Telecom) Connect() {
	log.Println("The telecom has been turned on.")
}

func (t Telecom) Disconnect() {
	log.Println("The telecom has been turned off.")
}
