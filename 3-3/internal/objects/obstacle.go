package objects

import "3-3/internal/contract"

type Obstacle struct {
	GameObject
}

func NewObstacle(gameMap *contract.Map) *Obstacle {
	return &Obstacle{
		GameObject: GameObject{
			gameMap: gameMap,
			name:    contract.ObstacleName,
		},
	}
}

func (o *Obstacle) Symbol() rune {
	return contract.ObstacleSymbol
}
