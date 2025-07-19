package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

// ShiftService định nghĩa các phương thức dịch vụ cho Shift
type ShiftService struct {
	repo repository.ShiftRepo
}

// NewShiftService tạo một thể hiện mới của ShiftService
func NewShiftService(r repository.ShiftRepo) *ShiftService {
	return &ShiftService{repo: r}
}

func (s *ShiftService) GetAll(page, pageSize int) ([]model.Shift, int64, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *ShiftService) GetByID(id uint) (*model.Shift, error) {
	return s.repo.FindByID(id)
}

func (s *ShiftService) Create(shift *model.Shift) error {
	return s.repo.Create(shift)
}

func (s *ShiftService) Update(id uint, updates map[string]interface{}) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Update(id, updates)
}

func (s *ShiftService) Delete(id uint) error {
	return s.repo.Delete(id)
}
