package server

type MessageRequest struct {
	Message string `json:"message" binding:"required"`
}
