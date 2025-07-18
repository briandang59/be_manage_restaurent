package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type CustomerService struct {
	repo repository.CustomerRepo
}

func NewCustomerService(r repository.CustomerRepo) *CustomerService {
	return &CustomerService{repo: r}
}

func (s *CustomerService) GetAll(page, pageSize int, preloadFields []string) ([]model.Customer, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}
