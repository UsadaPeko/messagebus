package messegebus_test

import (
	"github.com/stretchr/testify/assert"
	messegebus "gopkg.in/UsadaPeko/messaegbus"
	"testing"
)

func Test_PubSub_EmptyChan(t *testing.T) {
	bus := messegebus.New()

	assert.NotPanics(t, func() {
		bus.Pub(
			messegebus.Message{
				ID:       "",
				Metadata: "",
			},
		)
	})
}

func Test_PubSub(t *testing.T) {
	message := messegebus.Message{
		ID:       "",
		Metadata: "",
	}
	done := make(chan struct{})

	channel := make(chan messegebus.Message)
	go func(channel <-chan messegebus.Message) {
		msg := <-channel
		assert.Equal(t, message.ID, msg.ID)
		done <- struct{}{}
	}(channel)

	bus := messegebus.New(channel)

	bus.Pub(message)

	<-done
}

func Test_PubSub_EmptySub(t *testing.T) {
	message := messegebus.Message{
		ID:       "",
		Metadata: "",
	}
	done := make(chan struct{})

	// First Sub
	channel1 := make(chan messegebus.Message)
	go func(channel <-chan messegebus.Message) {
		msg := <-channel
		assert.Equal(t, message.ID, msg.ID)
		done <- struct{}{}
	}(channel1)
	// Second Sub
	channel2 := make(chan messegebus.Message)
	go func(channel <-chan messegebus.Message) {
		msg := <-channel
		assert.Equal(t, message.ID, msg.ID)
		done <- struct{}{}
	}(channel2)

	bus := messegebus.New(channel1, channel2)

	bus.Pub(message)

	<-done
	<-done
}

func Test_PubSub_MultiSub(t *testing.T) {
	channel := make(chan messegebus.Message)
	message := messegebus.Message{
		ID:       "",
		Metadata: "",
	}

	bus := messegebus.New(channel)

	bus.Pub(message)
}
