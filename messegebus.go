package messegebus

import "sync"

type Bus struct {
	channels []chan<- Message

	mutex sync.RWMutex
}

func New(channels ...chan<- Message) *Bus {
	return &Bus{
		channels: channels,
	}
}

type Message struct {
	ID       string
	Metadata string
}

func (b *Bus) Pub(message Message) {
	go func() {
		for i := range b.channels {
			b.channels[i] <- message
		}
	}()
}

func (b *Bus) Sub(channel chan<- Message) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.channels = append(b.channels, channel)
}
