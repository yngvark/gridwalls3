package mainhelp_test

import (
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/log2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/mainhelp"
)

func TestOslookup(t *testing.T) {

	t.Run("Should parse cors worigins", func(t *testing.T) {
		lg, err := log2.New()
		assert.Nil(t, err)

		m := mainhelp.New(lg)
		allowed, err := m.GetAllowedCorsOrigins(osLookupEnv, "TEST_ENV")
		assert.Nil(t, err)

		expected := make(map[string]bool)
		expected["http://localhost:3000"] = true
		expected["https://localhost:3001"] = true

		assert.Equal(t, expected, allowed)
	})
}

func osLookupEnv(_ string) (string, bool) {
	return "http://localhost:3000,https://localhost:3001", true
}
