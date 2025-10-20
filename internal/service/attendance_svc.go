package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
	"time"
)

type AttendanceService struct {
	repo *repository.AttendanceRepo
}

func NewAttendanceService(r *repository.AttendanceRepo) *AttendanceService {
	return &AttendanceService{repo: r}
}

func (s *AttendanceService) Create(attendance *model.Attendance) error {
	// Tự động tính toán hours từ actual_start_time và actual_end_time
	if !attendance.ActualStartTime.IsZero() && !attendance.ActualEndTime.IsZero() {
		duration := attendance.ActualEndTime.Sub(attendance.ActualStartTime)
		attendance.Hours = int64(duration.Hours())
	}
	return s.repo.Create(attendance)
}

func (s *AttendanceService) GetByID(id uint) (*model.Attendance, error) {
	return s.repo.GetByID(id)
}

func (s *AttendanceService) Update(id uint, updates map[string]interface{}) error {
	// Nếu cập nhật actual_start_time hoặc actual_end_time, tự động tính lại hours
	if startTime, hasStartTime := updates["actual_start_time"]; hasStartTime {
		if endTime, hasEndTime := updates["actual_end_time"]; hasEndTime {
			if start, ok := startTime.(time.Time); ok {
				if end, ok := endTime.(time.Time); ok {
					if !start.IsZero() && !end.IsZero() {
						duration := end.Sub(start)
						updates["hours"] = int64(duration.Hours())
					}
				}
			}
		}
	}
	return s.repo.Update(id, updates)
}

func (s *AttendanceService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AttendanceService) List(employeeID *uint, offset, limit int) ([]model.Attendance, int64, error) {
	return s.repo.List(employeeID, offset, limit)
}
