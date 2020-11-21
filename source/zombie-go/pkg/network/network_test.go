package network_test

import (
	"testing"
	net "zombie-go/pkg/network"

	"github.com/stretchr/testify/assert"
)

func TestNetwork(t *testing.T) {
	t.Run("Should send message to listeners", func(t *testing.T) {
		// Given
		network := net.NewBroadcaster()
		testReceiver := &testReceiver{}

		network.AddListener(testReceiver)

		// When
		network.NotifyListeneres("YO")

		// Then
		assert.Equal(t, "YO", testReceiver.lastMsgReceived)
	})
}

type testReceiver struct {
	lastMsgReceived string
}

func (t *testReceiver) NotifyListeneres(msg string) {
	t.lastMsgReceived = msg
}
