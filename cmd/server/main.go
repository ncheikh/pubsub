package main

import (
	"github.com/rs/zerolog"

	"github.com/ncheikh/pubsub/internal/app/server"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	server := server.New()

	server.Start()
}
