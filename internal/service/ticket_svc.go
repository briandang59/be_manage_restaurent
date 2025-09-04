package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type TicketService struct {
	repo *repository.TicketRepo
}

func NewTicketService(r *repository.TicketRepo) *TicketService {
	return &TicketService{repo: r}
}

func (s *TicketService) Create(ticket *model.Ticket) error {
	return s.repo.Create(ticket)
}

func (s *TicketService) GetByID(id uint, preloadFields []string) (*model.Ticket, error) {
	return s.repo.GetByID(id, preloadFields)
}

func (s *TicketService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *TicketService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *TicketService) List(offset, limit int, preloadFields []string) ([]model.Ticket, int64, error) {
	return s.repo.List(offset, limit, preloadFields)
}
