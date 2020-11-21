package zombie

import (
	"fmt"
	"math/rand"
	"zombie-go/pkg/worldmap"
)

type zombie struct {
	x        int
	y        int
	worldMap *worldmap.WorldMap
}

type zombieMove struct {
	X int
	Y int
}

func newZombie(x int, y int, worldMap *worldmap.WorldMap) *zombie {
	return &zombie{
		x:        x,
		y:        y,
		worldMap: worldMap,
	}
}

func newZombieMove(x int, y int) *zombieMove {
	return &zombieMove{
		X: x,
		Y: y,
	}
}

func (z *zombie) move() (*zombie, *zombieMove, error) {
	newX, err := z.getNewCoordPart(z.x, worldmap.Axis.X)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get new x coordinate: %w", err)
	}

	newY, err := z.getNewCoordPart(z.y, worldmap.Axis.Y)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get new x coordinate: %w", err)
	}

	newZ := newZombie(newX, newY, z.worldMap)
	move := newZombieMove(newX, newY)

	return newZ, move, nil
}

func (z *zombie) getNewCoordPart(currentValue int, axisType worldmap.AxisType) (int, error) {
	direction := rand.Intn(3) - 1 //nolint:gomnd,gosec    // [-1, 1]
	suggestion := currentValue + direction

	var newValue int

	isInMap, err := z.worldMap.IsInMap(suggestion, axisType)
	if err != nil {
		return -1, fmt.Errorf("could not detect if value is within map: %w", err)
	}

	if isInMap {
		newValue = suggestion
	} else {
		newValue = currentValue
	}

	return newValue, nil
}
