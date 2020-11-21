package zombie_test

import (
	"math/rand"
	"testing"

	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/worldmap"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/zombie"

	"github.com/stretchr/testify/assert"
)

func TestZombie(t *testing.T) {
	t.Run("Should move randomly", func(t *testing.T) {
		// Given
		m := worldmap.New(20, 10)                                     //nolint:gomnd
		z := zombie.NewZombie(10, 5, m, rand.New(rand.NewSource(45))) //nolint:gosec,gomnd
		g := zombie.NewGenerator(z)

		// When+Then
		assertNextPosition(t, g, 9, 5)
		assertNextPosition(t, g, 8, 4)
		assertNextPosition(t, g, 9, 4)
		assertNextPosition(t, g, 8, 5)
		assertNextPosition(t, g, 8, 5)
	})
}

func assertNextPosition(t *testing.T, generator *zombie.Generator, x int, y int) {
	move, err := generator.Next()
	assert.Nil(t, err)
	assert.Equal(t, zombie.NewZombieMove(x, y), move)
}
