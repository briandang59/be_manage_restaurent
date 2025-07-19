package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type AvailibilityRepo interface {
	FindAll(page, pageSize int, preloadFields []string) ([]model.Availibility, int64, error)
	FindByID(id uint) (*model.Availibility, error)
	Create(availibility *model.Availibility) error
	BulkCreate(availibilities []model.Availibility) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type availibilityRepo struct {
	db *gorm.DB
}

func NewAvailibilityRepo(db *gorm.DB) AvailibilityRepo {
	return &availibilityRepo{db: db}
}

func (r *availibilityRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Availibility, int64, error) {
	var list []model.Availibility
	var total int64
	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Availibility{})
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

func (r *availibilityRepo) FindByID(id uint) (*model.Availibility, error) {
	var availibility model.Availibility
	if err := r.db.First(&availibility, id).Error; err != nil {
		return nil, err
	}
	return &availibility, nil
}

func (r *availibilityRepo) Create(availibility *model.Availibility) error {
	return r.db.Create(availibility).Error
}

func (r *availibilityRepo) BulkCreate(availibilities []model.Availibility) error {
	return r.db.Create(&availibilities).Error
}

func (r *availibilityRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Availibility{}).Where("id = ?", id).Updates(updates).Error
}

func (r *availibilityRepo) Delete(id uint) error {
	return r.db.Delete(&model.Availibility{}, id).Error
}
