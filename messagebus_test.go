package messagebus_test

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/UsadaPeko/messagbus"
	"testing"
)

func Test_PubSub_EmptyChan(t *testing.T) {
	bus := messagebus.New()

	assert.NotPanics(t, func() {
		bus.Pub(
			messagebus.Message{
				ID:       "",
				Metadata: "",
			},
		)
	})
}

func Test_PubSub(t *testing.T) {
	message := messagebus.Message{
		ID:       "",
		Metadata: "",
	}
	done := make(chan struct{})

	channel := make(chan messagebus.Message)
	go func(channel <-chan messagebus.Message) {
		msg := <-channel
		assert.Equal(t, message.ID, msg.ID)
		done <- struct{}{}
	}(channel)

	bus := messagebus.New(channel)

	bus.Pub(message)

	<-done
}

func Test_PubSub_MultiSub(t *testing.T) {
	message := messagebus.Message{
		ID:       "",
		Metadata: "",
	}
	done := make(chan struct{})

	// First Sub
	channel1 := make(chan messagebus.Message)
	go func(channel <-chan messagebus.Message) {
		msg := <-channel
		assert.Equal(t, message.ID, msg.ID)
		done <- struct{}{}
	}(channel1)
	// Second Sub
	channel2 := make(chan messagebus.Message)
	go func(channel <-chan messagebus.Message) {
		msg := <-channel
		assert.Equal(t, message.ID, msg.ID)
		done <- struct{}{}
	}(channel2)

	bus := messagebus.New(channel1, channel2)

	bus.Pub(message)

	<-done
	<-done
}

func Test_PubSub_EmptySub(t *testing.T) {
	channel := make(chan messagebus.Message)
	message := messagebus.Message{
		ID:       "",
		Metadata: "",
	}

	bus := messagebus.New(channel)

	bus.Pub(message)
}
