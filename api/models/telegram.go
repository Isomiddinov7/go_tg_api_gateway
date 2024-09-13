package models

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type MessageResponse struct {
	ChatID  string `json:"chat_id"`
	Message string `json:"message"`
	File    string `json:"file"`
}
