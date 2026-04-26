package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"umkm-chatbot/internal/model"
)

// TelegramClient handles communication with the Telegram Bot API
type TelegramClient interface {
	SendMessage(chatID int64, text string) error
}

type telegramClient struct {
	botToken string
}

func NewTelegramClient(botToken string) TelegramClient {
	return &telegramClient{botToken: botToken}
}

// SendMessage sends a text message to a specific Telegram chat
func (c *telegramClient) SendMessage(chatID int64, text string) error {
	// Construct the Telegram API URL with the bot token
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.botToken)

	// Prepare the request payload
	reqBody := model.SendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Perform the POST request to Telegram API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-200 responses and capture the error message from Telegram
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram api error: %s (status code: %d), body: %s", resp.Status, resp.StatusCode, string(body))
	}

	return nil
}
