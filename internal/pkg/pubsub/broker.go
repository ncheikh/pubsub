package pubsub

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type publishMessage struct {
	Message string
}

// Broker represents a broker
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
			// Consume multiple channels and react.  The use of multiple channels removes
			// several race conditions as each action will block
			// TODO test what happens if a client disconnects in the middle of a message stream
			select {
			case event := <-b.publish:
				// This will block on each client, and thus it is important to remove clients
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

// RegisterClient in registry and make available
func (b *Broker) RegisterClient(client Client) error {
	if !(b.active) {
		return errors.New("Broker Not Started")
	}

	b.connect <- client

	return nil
}

// UnregisterClient removes client from client registry
func (b *Broker) UnregisterClient(client Client) error {
	if !(b.active) {
		return errors.New("Broker Not Started")
	}

	b.disconnect <- client

	return nil
}

// PublishMessage to clients
func (b *Broker) PublishMessage(m Message) error {
	if !(b.active) {
		return errors.New("Broker Not Started")
	}

	b.publish <- publishMessage{
		Message: m.Message,
	}

	return nil
}
