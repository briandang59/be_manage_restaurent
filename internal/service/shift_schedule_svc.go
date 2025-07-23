package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

// ShiftScheduleService định nghĩa các phương thức dịch vụ cho ShiftSchedule
type ShiftScheduleService struct {
	repo repository.ShiftScheduleRepo
}

// NewShiftScheduleService tạo một thể hiện mới của ShiftScheduleService
func NewShiftScheduleService(r repository.ShiftScheduleRepo) *ShiftScheduleService {
	return &ShiftScheduleService{repo: r}
}

func (s *ShiftScheduleService) GetAll(page, pageSize int, preloadFields []string, filters map[string]interface{}) ([]model.ShiftSchedule, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields, filters)
}

func (s *ShiftScheduleService) GetByID(id uint) (*model.ShiftSchedule, error) {
	return s.repo.FindByID(id)
}

func (s *ShiftScheduleService) Create(shiftSchedule *model.ShiftSchedule) error {
	return s.repo.Create(shiftSchedule)
}

func (s *ShiftScheduleService) BulkCreate(shiftSchedules []model.ShiftSchedule) error {
	return s.repo.BulkCreate(shiftSchedules)
}

func (s *ShiftScheduleService) Update(id uint, updates map[string]interface{}) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Update(id, updates)
}

func (s *ShiftScheduleService) Delete(id uint) error {
	return s.repo.Delete(id)
}
