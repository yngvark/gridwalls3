package worldmap

import (
	"fmt"
)

type AxisType int

var Axis = struct { //nolint:gochecknoglobals
	X AxisType
	Y AxisType
}{
	X: 1,
	Y: 2, //nolint:gomnd
}

type WorldMap struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func New(maxX int, maxY int) *WorldMap {
	return &WorldMap{
		MinX: 0,
		MaxX: 0,
		MinY: maxX,
		MaxY: maxY,
	}
}

func (m *WorldMap) IsInMap(axisValue int, axis AxisType) (bool, error) {
	switch axis {
	case Axis.X:
		return m.XIsInMap(axisValue), nil
	case Axis.Y:
		return m.YIsInMap(axisValue), nil
	default:
		return false, fmt.Errorf("not a valid Axis type: %d", axis)
	}
}

func (m *WorldMap) XIsInMap(x int) bool {
	return x <= m.MaxX && x >= m.MinX
}

func (m *WorldMap) YIsInMap(y int) bool {
	return y <= m.MaxY && y >= m.MinY
}
