package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepo
}

func NewOrderService(r *repository.OrderRepo) *OrderService {
	return &OrderService{repo: r}
}

func (s *OrderService) Create(order *model.Order) error {
	return s.repo.Create(order)
}

func (s *OrderService) GetByID(id uint) (*model.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *OrderService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *OrderService) List(offset, limit int) ([]model.Order, int64, error) {
	return s.repo.List(offset, limit)
}

func (s *OrderService) FindOrderByTable(id uint) (*model.Order, error) {
	return s.repo.FindOrderByTableId(id)
}
