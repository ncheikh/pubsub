package pubsub

import uuid "github.com/satori/go.uuid"

type Client struct {
	ID      string
	Channel chan string
}

func NewClient() Client {
	return Client{
		ID:      uuid.NewV4().String(),
		Channel: make(chan string),
	}
}
