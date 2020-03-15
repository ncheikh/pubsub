package server

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ncheikh/pubsub/internal/pkg/pubsub"
)

func (s *Server) handlePublish(c *gin.Context) {
	var messageRequest MessageRequest

	if err := c.ShouldBindJSON(&messageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := pubsub.Message{Message: messageRequest.Message}

	if err := s.broker.PublishMessage(msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message published"})
}

func (s *Server) handleSubscribe(c *gin.Context) {

	client := pubsub.NewClient()

	s.broker.RegisterClient(client)

	go func() {
		select {
		case <-c.Request.Context().Done():
			s.broker.UnregisterClient(client)
		}
	}()

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", <-client.Channel)
		return true
	})
}
