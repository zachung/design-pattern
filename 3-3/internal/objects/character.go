package objects

import (
	"3-3/internal/contract"
	"3-3/internal/states"
	"math/rand/v2"
)

type Character struct {
	RoleStruct
	direction rune
}

func NewCharacter(gameMap *contract.Map) *Character {
	character := &Character{
		RoleStruct: RoleStruct{
			GameObject: GameObject{
				gameMap: gameMap,
				name:    contract.CharacterName,
			},
			health:    300,
			maxHealth: 300,
		},
		direction: randDirection(),
	}
	role := contract.Role(character)
	character.State = *states.NewNormal(&role)

	return character
}

func randDirection() rune {
	m := []rune{contract.CharacterUpSymbol, contract.CharacterRightSymbol, contract.CharacterDownSymbol, contract.CharacterLeftSymbol}

	return m[rand.IntN(4)]
}

func (c *Character) Symbol() rune {
	return c.direction
}

func (c *Character) SetDirection(direction rune) {
	c.direction = direction
}
