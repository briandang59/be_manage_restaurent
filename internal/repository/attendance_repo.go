package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type AttendanceRepo struct {
	db *gorm.DB
}

func NewAttendanceRepo(db *gorm.DB) *AttendanceRepo {
	return &AttendanceRepo{db: db}
}

func (r *AttendanceRepo) Create(attendance *model.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *AttendanceRepo) GetByID(id uint) (*model.Attendance, error) {
	var attendance model.Attendance
	if err := r.db.Preload("ShiftSchedule").First(&attendance, id).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *AttendanceRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Attendance{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AttendanceRepo) Delete(id uint) error {
	return r.db.Delete(&model.Attendance{}, id).Error
}

func (r *AttendanceRepo) List(offset, limit int) ([]model.Attendance, int64, error) {
	var attendances []model.Attendance
	var total int64
	r.db.Model(&model.Attendance{}).Count(&total)
	if err := r.db.Preload("ShiftSchedule").Offset(offset).Limit(limit).Find(&attendances).Error; err != nil {
		return nil, 0, err
	}
	return attendances, total, nil
} 