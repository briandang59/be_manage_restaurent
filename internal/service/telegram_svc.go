package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type TelegramService struct {
	repo *repository.TelegramRepo
}

func NewTelegramService(r *repository.TelegramRepo) *TelegramService {
	return &TelegramService{repo: r}
}

func (s *TelegramService) SendMessage(req model.TelegramSendRequest) (*model.TelegramSendResponse, error) {
	return s.repo.Send(req)
}
