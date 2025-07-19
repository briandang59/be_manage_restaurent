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

// Thêm phương thức mới cho CRUD
func (s *CustomerService) Create(customer *model.Customer) error {
	return s.repo.Create(customer)
}

func (s *CustomerService) Update(id uint, updatedCustomer *model.Customer) error {
	customer, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	// Áp dụng các cập nhật
	customer.FullName = updatedCustomer.FullName
	customer.PhoneNumber = updatedCustomer.PhoneNumber

	return s.repo.Update(customer)
}

func (s *CustomerService) Delete(id uint) error {
	return s.repo.Delete(id)
}
