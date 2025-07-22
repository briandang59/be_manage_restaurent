package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepo) GetByID(id uint) (*model.Order, error) {
	var order model.Order
	if err := r.db.Preload("Customer").Preload("Table").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Order{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OrderRepo) Delete(id uint) error {
	return r.db.Delete(&model.Order{}, id).Error
}

func (r *OrderRepo) List(offset, limit int) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64
	r.db.Model(&model.Order{}).Count(&total)
	if err := r.db.Preload("Customer").Preload("Table").Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
} 