package pubsub

import (
	"testing"
)

func TestRegisterClientOnInactiveBroker(t *testing.T) {
	broker := NewBroker()

	client := NewClient()

	err := broker.RegisterClient(client)

	if err == nil {
		t.Error("Expected Error")
	}
}

func TestUnregisterClientOnInactiveBroker(t *testing.T) {
	broker := NewBroker()

	client := NewClient()

	err := broker.UnregisterClient(client)

	if err == nil {
		t.Error("Expected Error")
	}
}

func TestPublishMessageOnInactiveBroker(t *testing.T) {
	broker := NewBroker()

	message := Message{}

	err := broker.PublishMessage(message)

	if err == nil {
		t.Error("Expected Error")
	}
}

func TestActiveBroker(t *testing.T) {
	broker := NewBroker()

	client := NewClient()

	message := Message{
		Message: "test",
	}

	broker.Start()

	var err error

	err = broker.RegisterClient(client)

	if err != nil {
		t.Error("No Error Expected", err)
	}

	err = broker.PublishMessage(message)

	if err != nil {
		t.Error("No Error Expected", err)
	}

	result := <-client.Channel

	if result != "test" {
		t.Error("Error expected client message to be", "test")
	}
}

func TestActiveBrokerWithMultipleSubscribers(t *testing.T) {
	broker := NewBroker()

	var clients []Client

	clientCount := 10

	for i := 0; i < clientCount; i++ {
		clients = append(clients, NewClient())
	}

	message := Message{
		Message: "test",
	}

	broker.Start()

	var err error

	for i := 0; i < clientCount; i++ {
		err = broker.RegisterClient(clients[i])

		if err != nil {
			t.Error("No Error Expected", err)
		}
	}

	results := []string{}

	for i := 0; i < clientCount; i++ {
		go func(client Client) {
			for {
				result := <-client.Channel

				results = append(results, result)

				if result != "test" {
					t.Error("Error expected client message to be", "test")
					t.FailNow()
				}
			}
		}(clients[i])
	}

	err = broker.PublishMessage(message)

	if err != nil {
		t.Error("No Error Expected", err)
	}

}
