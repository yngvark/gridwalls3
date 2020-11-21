package zombie

import (
	"fmt"
)

type Generator struct {
	zombie *Zombie
}

func NewGenerator(initialZombie *Zombie) *Generator {
	return &Generator{
		zombie: initialZombie,
	}
}

func (g *Generator) Next() (*Move, error) {
	z, move, err := g.zombie.move()
	if err != nil {
		return nil, fmt.Errorf("could not move zombie: %w", err)
	}

	g.zombie = z

	return move, nil
}
