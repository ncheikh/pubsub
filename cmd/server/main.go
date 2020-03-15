package main

import (
	"github.com/ncheikh/pubsub/internal/app/server"
)

func main() {
	initLogging()

	server := server.New()

	server.Start()
}
