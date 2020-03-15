package pubsub

import uuid "github.com/satori/go.uuid"

// Client represents a client that is subscribed
type Client struct {
	ID      string
	Channel chan string
}

// NewClient creates a new client
func NewClient() Client {
	return Client{
		ID:      uuid.NewV4().String(),
		Channel: make(chan string),
	}
}
