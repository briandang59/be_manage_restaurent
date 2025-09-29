package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type BookingRepo interface {
	FindAll(page, pageSize int, preloadFields []string) ([]model.Booking, int64, error)
	Create(booking *model.Booking) error
	FindByID(id uint) (*model.Booking, error)
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type bookingRepo struct {
	db *gorm.DB
}

func NewBookingRepo(db *gorm.DB) BookingRepo {
	return &bookingRepo{db: db}
}

func (r *bookingRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Booking, int64, error) {
	var list []model.Booking
	var total int64
	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Booking{})

	for _, field := range preloadFields {
		query = query.Preload(field)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *bookingRepo) Create(booking *model.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepo) FindByID(id uint) (*model.Booking, error) {
	var booking model.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepo) Delete(id uint) error {
	return r.db.Delete(&model.Booking{}, id).Error
}

func (r *bookingRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Booking{}).Where("id = ?", id).Updates(updates).Error
}
