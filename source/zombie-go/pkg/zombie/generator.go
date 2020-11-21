package zombie

import (
	"encoding/json"
	"fmt"
)

type generator struct {
	zombie *zombie
}

func newGenerator(z *zombie) *generator {
	return &generator{
		zombie: z,
	}
}

func (g *generator) next() (string, error) {
	z, move, err := g.zombie.move()
	if err != nil {
		return "", fmt.Errorf("could not move zombie: %w", err)
	}

	g.zombie = z

	jsn, err := json.Marshal(move)
	if err != nil {
		return "", fmt.Errorf("could not marshal zombie move: %w", err)
	}

	return string(jsn), nil
}
