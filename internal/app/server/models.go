package server

// MessageRequest represents a message request
type MessageRequest struct {
	Message string `json:"message" binding:"required"`
}
