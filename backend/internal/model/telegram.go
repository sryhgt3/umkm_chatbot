package model

// Update represents a Telegram update
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  *Message `json:"message"`
}

// Message represents a Telegram message
type Message struct {
	MessageID int          `json:"message_id"`
	From      TelegramUser `json:"from"`
	Chat      Chat         `json:"chat"`
	Text      string       `json:"text"`
}

// TelegramUser represents a Telegram user
type TelegramUser struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

// Chat represents a Telegram chat
type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

// SendMessageRequest represents the payload for sendMessage API
type SendMessageRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
