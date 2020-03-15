package pubsub

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type publishMessage struct {
	Message string
}

// Broker model
type Broker struct {
	active     bool
	publish    chan publishMessage
	connect    chan Client
	disconnect chan Client
	clients    map[string]Client
}

// NewBroker creates a new broker
func NewBroker() *Broker {
	return &Broker{
		active:     false,
		publish:    make(chan publishMessage),
		connect:    make(chan Client),
		disconnect: make(chan Client),
		clients:    map[string]Client{},
	}
}

// Start the broker running, we use channels to avoid race conditions.
func (b *Broker) Start() {
	b.active = true
	go func() {
		for {
			select {
			case event := <-b.publish:
				for id, client := range b.clients {
					log.Debug().Msgf("Publishing to: %s\n", id)

					client.Channel <- event.Message
				}

			case client := <-b.connect:
				b.clients[client.ID] = client

				log.Debug().Msgf("Registering Client: %s\n", client.ID)

			case client := <-b.disconnect:
				delete(b.clients, client.ID)

				log.Debug().Msgf("Unregistering Client: %s\n", client.ID)
			}
		}
	}()
}

func (b *Broker) RegisterClient(client Client) error {
	if !(b.active) {
		return errors.New("Broker Not Started")
	}

	b.connect <- client

	return nil
}

func (b *Broker) UnregisterClient(client Client) error {
	if !(b.active) {
		return errors.New("Broker Not Started")
	}

	b.disconnect <- client

	return nil
}

func (b *Broker) PublishMessage(m Message) error {
	if !(b.active) {
		return errors.New("Broker Not Started")
	}

	b.publish <- publishMessage{
		Message: m.Message,
	}

	return nil
}
