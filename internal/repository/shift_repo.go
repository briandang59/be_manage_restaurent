package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

// ShiftRepo định nghĩa các phương thức tương tác với bảng Shift
type ShiftRepo interface {
	FindAll(page, pageSize int) ([]model.Shift, int64, error)
	FindByID(id uint) (*model.Shift, error)
	Create(shift *model.Shift) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type shiftRepo struct {
	db *gorm.DB
}

// NewShiftRepo tạo một thể hiện mới của ShiftRepo
func NewShiftRepo(db *gorm.DB) ShiftRepo {
	return &shiftRepo{db: db}
}

func (r *shiftRepo) FindAll(page, pageSize int) ([]model.Shift, int64, error) {
	var list []model.Shift
	var total int64
	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.Shift{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *shiftRepo) FindByID(id uint) (*model.Shift, error) {
	var shift model.Shift
	if err := r.db.First(&shift, id).Error; err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *shiftRepo) Create(shift *model.Shift) error {
	return r.db.Create(shift).Error
}

func (r *shiftRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Shift{}).Where("id = ?", id).Updates(updates).Error
}
func (r *shiftRepo) Delete(id uint) error {
	return r.db.Delete(&model.Shift{}, id).Error
}
