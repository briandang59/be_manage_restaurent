package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

// EmployeeService định nghĩa các phương thức dịch vụ cho Employee
type EmployeeService struct {
	repo repository.EmployeeRepo
}

// NewEmployeeService tạo một thể hiện mới của EmployeeService
func NewEmployeeService(r repository.EmployeeRepo) *EmployeeService {
	return &EmployeeService{repo: r}
}

func (s *EmployeeService) GetAll(page, pageSize int, preloadFields []string) ([]model.Employee, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}

func (s *EmployeeService) GetByID(id uint) (*model.Employee, error) {
	return s.repo.FindByID(id)
}

func (s *EmployeeService) Create(employee *model.Employee) error {
	return s.repo.Create(employee)
}

// Phương thức Update mới cho phép cập nhật một phần
func (s *EmployeeService) Update(id uint, updates map[string]interface{}) error {
	// Lấy bản ghi cũ để đảm bảo nó tồn tại, sau đó cập nhật
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Update(id, updates)
}

func (s *EmployeeService) Delete(id uint) error {
	return s.repo.Delete(id)
}
