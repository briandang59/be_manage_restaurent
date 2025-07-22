package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type OrderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepo(db *gorm.DB) *OrderItemRepo {
	return &OrderItemRepo{db: db}
}

func (r *OrderItemRepo) Create(orderItem *model.OrderItem) error {
	return r.db.Create(orderItem).Error
}

func (r *OrderItemRepo) GetByID(id uint) (*model.OrderItem, error) {
	var orderItem model.OrderItem
	if err := r.db.Preload("Order").First(&orderItem, id).Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (r *OrderItemRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.OrderItem{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OrderItemRepo) Delete(id uint) error {
	return r.db.Delete(&model.OrderItem{}, id).Error
}

func (r *OrderItemRepo) List(offset, limit int) ([]model.OrderItem, int64, error) {
	var orderItems []model.OrderItem
	var total int64
	r.db.Model(&model.OrderItem{}).Count(&total)
	if err := r.db.Preload("Order").Offset(offset).Limit(limit).Find(&orderItems).Error; err != nil {
		return nil, 0, err
	}
	return orderItems, total, nil
} 