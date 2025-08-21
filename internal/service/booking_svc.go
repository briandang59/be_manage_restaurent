package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type BookingService struct {
	repo repository.BookingRepo
}

func NewBookingService(r repository.BookingRepo) *BookingService {
	return &BookingService{repo: r}
}

func (s *BookingService) Create(booking *model.Booking) error {
	return s.repo.Create(booking)
}

func (s *BookingService) GetByID(id uint) (*model.Booking, error) {
	return s.repo.FindByID(id)
}

func (s *BookingService) Update(booking *model.Booking) error {
	return s.repo.Update(booking)
}

func (s *BookingService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *BookingService) List(page, pageSize int, preloadFields []string) ([]model.Booking, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}
