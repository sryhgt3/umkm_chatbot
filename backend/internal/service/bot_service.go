package service

import (
	"strings"
	"umkm-chatbot/internal/client"
	"umkm-chatbot/internal/model"
)

// BotService defines the interface for processing bot updates
type BotService interface {
	HandleUpdate(update model.Update) error
}

type botService struct {
	tgClient client.TelegramClient
}

func NewBotService(tgClient client.TelegramClient) BotService {
	return &botService{tgClient: tgClient}
}

// HandleUpdate processes a single update from Telegram and decides the response logic
func (s *botService) HandleUpdate(update model.Update) error {
	// Only process text messages; ignore photos, voice, etc. for now
	if update.Message == nil || update.Message.Text == "" {
		return nil
	}

	chatID := update.Message.Chat.ID
	// Normalize the input text for easier matching
	text := strings.ToLower(strings.TrimSpace(update.Message.Text))

	// Simple Command Router: Matches text input to specific responses
	var reply string
	switch text {
	case "halo":
		reply = "Halo 👋 mau pesan apa?"
	case "menu":
		reply = "Menu:\n1. Nasi Goreng - 15000\n2. Mie Ayam - 12000\n3. Es Teh - 5000"
	default:
		// Default fallback message if no command is matched
		reply = "Ketik 'menu' untuk melihat daftar makanan"
	}

	// Send the reply back to the user via the Telegram Client
	return s.tgClient.SendMessage(chatID, reply)
}
