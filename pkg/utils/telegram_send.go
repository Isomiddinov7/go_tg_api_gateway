package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_tg_api_gateway/api/models"
	"go_tg_api_gateway/genproto/users_service"
	"net/http"
)

func SendMessageToTelegram(req *users_service.TelegramMessageResponse) error {
	token := "7135277002:AAG9phZbadRo-iafeN6fQDvMVPyrGtN-k3E"
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	msg := models.MessageResponse{
		ChatID:  req.TelegramId,
		Message: req.Message,
		File:    req.File,
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Telegram API error: %s", resp.Status)
	}

	return nil
}
