package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

// ShiftScheduleRepo định nghĩa các phương thức tương tác với bảng ShiftSchedule
type ShiftScheduleRepo interface {
	FindAll(page, pageSize int, preloadFields []string) ([]model.ShiftSchedule, int64, error)
	FindByID(id uint) (*model.ShiftSchedule, error)
	Create(shiftSchedule *model.ShiftSchedule) error
	BulkCreate(shiftSchedules []model.ShiftSchedule) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type shiftScheduleRepo struct {
	db *gorm.DB
}

// NewShiftScheduleRepo tạo một thể hiện mới của ShiftScheduleRepo
func NewShiftScheduleRepo(db *gorm.DB) ShiftScheduleRepo {
	return &shiftScheduleRepo{db: db}
}

func (r *shiftScheduleRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.ShiftSchedule, int64, error) {
	var list []model.ShiftSchedule
	var total int64
	offset := (page - 1) * pageSize

	query := r.db.Model(&model.ShiftSchedule{})
	for _, field := range preloadFields {
		query = query.Preload(field)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *shiftScheduleRepo) FindByID(id uint) (*model.ShiftSchedule, error) {
	var shiftSchedule model.ShiftSchedule
	if err := r.db.First(&shiftSchedule, id).Error; err != nil {
		return nil, err
	}
	return &shiftSchedule, nil
}

func (r *shiftScheduleRepo) Create(shiftSchedule *model.ShiftSchedule) error {
	return r.db.Create(shiftSchedule).Error
}

func (r *shiftScheduleRepo) BulkCreate(shiftSchedules []model.ShiftSchedule) error {
	return r.db.Create(&shiftSchedules).Error
}

func (r *shiftScheduleRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.ShiftSchedule{}).Where("id = ?", id).Updates(updates).Error
}

func (r *shiftScheduleRepo) Delete(id uint) error {
	return r.db.Delete(&model.ShiftSchedule{}, id).Error
}
