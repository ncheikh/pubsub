package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ncheikh/pubsub/internal/pkg/pubsub"
)

type PubSubBroker interface {
	Start()
	RegisterClient(pubsub.Client) error
	UnregisterClient(pubsub.Client) error

	PublishMessage(pubsub.Message) error
}

type Server struct {
	router *gin.Engine
	broker PubSubBroker
}

func New() *Server {
	// Create Gin Instance
	r := gin.New()
	r.Use(gin.Recovery())

	// Instantiate Broker
	broker := pubsub.NewBroker()

	return &Server{
		router: r,
		broker: broker,
	}
}

func (s *Server) Start() {
	// Bind Handlers
	s.bind()

	// Start Broker
	s.broker.Start()

	// Start Server
	s.router.Run()
}

func (s *Server) bind() {
	s.router.POST("/publish", s.handlePublish)

	s.router.GET("/subscribe", s.handleSubscribe)
}
