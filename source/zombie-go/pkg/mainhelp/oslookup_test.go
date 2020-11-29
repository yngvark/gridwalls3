package mainhelp_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/mainhelp"
	"testing"
)

func TestOslookup(t *testing.T) {
	t.Run("Should parse cors worigins", func(t *testing.T) {
		allowed, err := mainhelp.GetAllowedCorsOrigins(osLookupEnv, "TEST_ENV")
		assert.Nil(t, err)

		expected := make(map[string]bool)
		expected["http://localhost:3000"] = true
		expected["https://localhost:3001"] = true

		assert.Equal(t, expected, allowed)
	})
}

func osLookupEnv(key string) (string, bool) {
	return "http://localhost:3000,https://localhost:3001", true
}
