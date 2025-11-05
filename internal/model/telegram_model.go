package model

type TelegramSendRequest struct {
	BotToken string `json:"bot_token"`
	ChatID   string `json:"chat_id"`
	Message  string `json:"message"`
}

type TelegramSendResponse struct {
	OK     bool        `json:"ok"`
	Result interface{} `json:"result"`
}
