package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type CustomerRepo interface {
	FindAll(page, pageSize int, preloadFields []string) ([]model.Customer, int64, error)
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) CustomerRepo {
	return &customerRepo{db: db}
}

func (r *customerRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Customer, int64, error) {
	var list []model.Customer
	var total int64
	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Customer{})

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
