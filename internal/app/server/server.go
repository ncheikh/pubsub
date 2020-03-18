package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ncheikh/pubsub/internal/pkg/pubsub"
)

// pubSubBroker that allows for subscriptions and publishing
type pubSubBroker interface {
	Start()

	RegisterClient(pubsub.Client) error
	UnregisterClient(pubsub.Client) error

	PublishMessage(pubsub.Message) error
}

// Server object
type Server struct {
	router *gin.Engine
	broker pubSubBroker
}

// New creates a new server instance
func New() *Server {
	// Create Gin Instance
	r := gin.New()
	r.Use(gin.Recovery())

	// Instantiate Broker
	broker := pubsub.NewBroker()

	// HTTP/2 would be even better as it would use multiplexing
	return &Server{
		router: r,
		broker: broker,
	}
}

// Start the server listening
func (s *Server) Start() {
	// Bind Handlers
	s.bind()

	// Start Broker
	s.broker.Start()

	// Start Server
	s.router.Run()
}

// bind handlers to endpoints
func (s *Server) bind() {
	s.router.POST("/publish", s.handlePublish)

	s.router.GET("/subscribe", s.handleSubscribe)
}
