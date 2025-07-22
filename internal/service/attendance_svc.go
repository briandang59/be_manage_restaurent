package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type AttendanceService struct {
	repo *repository.AttendanceRepo
}

func NewAttendanceService(r *repository.AttendanceRepo) *AttendanceService {
	return &AttendanceService{repo: r}
}

func (s *AttendanceService) Create(attendance *model.Attendance) error {
	return s.repo.Create(attendance)
}

func (s *AttendanceService) GetByID(id uint) (*model.Attendance, error) {
	return s.repo.GetByID(id)
}

func (s *AttendanceService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *AttendanceService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AttendanceService) List(offset, limit int) ([]model.Attendance, int64, error) {
	return s.repo.List(offset, limit)
} 