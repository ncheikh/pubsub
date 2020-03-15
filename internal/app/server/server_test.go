package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestPublish(t *testing.T) {
	server := New()

	go func() {
		server.Start()
	}()

	time.Sleep(1 * time.Second)

	// Make Publish Request
	values := map[string]string{"message": "test"}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post("http://localhost:8080/publish", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Error("Expected response, got", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Expected 200 response, got", resp.StatusCode)
	}
}

// func TestSingleSubscriber(t *testing.T) {
// 	server := New()

// 	var resp *http.Response
// 	var err error

// 	go func() {
// 		server.Start()

// 		resp, err = http.Get("http://localhost:8080/subscribe")
// 		if err != nil {
// 			t.Error("err", err)

// 		}

// 		t.Log("resp", resp)
// 	}()

// 	time.Sleep(1 * time.Second)

// 	// Make Publish Request
// 	values := map[string]string{"message": "test"}

// 	jsonValue, _ := json.Marshal(values)

// 	resp, err = http.Post("http://localhost:8080/publish", "application/json", bytes.NewBuffer(jsonValue))

// 	t.Error("Expected 200 response, got", resp.StatusCode)

// 	if err != nil {
// 		t.Error("Expected response, got", err)
// 	}

// 	if resp.StatusCode != 200 {
// 		t.Error("Expected 200 response, got", resp.StatusCode)
// 	}
// }
