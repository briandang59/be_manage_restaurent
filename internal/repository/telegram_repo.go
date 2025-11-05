package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"manage_restaurent/internal/model"
	"net/http"
)

type TelegramRepo struct{}

func NewTelegramRepo() *TelegramRepo {
	return &TelegramRepo{}
}

func (r *TelegramRepo) Send(req model.TelegramSendRequest) (*model.TelegramSendResponse, error) {
	if req.BotToken == "" {
		return nil, errors.New("bot_token is required")
	}

	url := "https://api.telegram.org/bot" + req.BotToken + "/sendMessage"

	bodyData := map[string]interface{}{
		"chat_id": req.ChatID,
		"text":    req.Message,
	}

	body, _ := json.Marshal(bodyData)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.TelegramSendResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
