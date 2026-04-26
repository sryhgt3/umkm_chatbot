package handler

import (
	"log"
	"net/http"
	"umkm-chatbot/internal/model"
	"umkm-chatbot/internal/service"

	"github.com/gin-gonic/gin"
)

// WebhookHandler responsible for handling incoming HTTP requests from Telegram
type WebhookHandler struct {
	botService service.BotService
}

func NewWebhookHandler(botService service.BotService) *WebhookHandler {
	return &WebhookHandler{botService: botService}
}

// HandleTelegram processes POST requests sent by the Telegram Bot API
func (h *WebhookHandler) HandleTelegram(c *gin.Context) {
	var update model.Update
	// Unmarshal the incoming JSON into the Telegram Update model
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Delegate the business logic to the BotService
	if err := h.botService.HandleUpdate(update); err != nil {
		// Log the error for server-side debugging
		log.Printf("Failed to process update: %v", err)
		// Return 500 so Telegram knows the processing failed (and might retry)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process update", "details": err.Error()})
		return
	}

	// Acknowledge receipt of the update with a 200 OK
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
