package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type OrderItemService struct {
	repo *repository.OrderItemRepo
}

func NewOrderItemService(r *repository.OrderItemRepo) *OrderItemService {
	return &OrderItemService{repo: r}
}

func (s *OrderItemService) Create(orderItem *model.OrderItem) error {
	return s.repo.Create(orderItem)
}

func (s *OrderItemService) GetByID(id uint) (*model.OrderItem, error) {
	return s.repo.GetByID(id)
}

func (s *OrderItemService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *OrderItemService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *OrderItemService) List(offset, limit int) ([]model.OrderItem, int64, error) {
	return s.repo.List(offset, limit)
} 