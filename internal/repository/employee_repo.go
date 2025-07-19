package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

// EmployeeRepo định nghĩa các phương thức tương tác với bảng Employee
type EmployeeRepo interface {
	FindAll(page, pageSize int, preloadFields []string) ([]model.Employee, int64, error)
	FindByID(id uint) (*model.Employee, error)
	Create(employee *model.Employee) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type employeeRepo struct {
	db *gorm.DB
}

// NewEmployeeRepo tạo một thể hiện mới của EmployeeRepo
func NewEmployeeRepo(db *gorm.DB) EmployeeRepo {
	return &employeeRepo{db: db}
}

func (r *employeeRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Employee, int64, error) {
	var list []model.Employee
	var total int64
	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Employee{})
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

func (r *employeeRepo) FindByID(id uint) (*model.Employee, error) {
	var employee model.Employee
	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepo) Create(employee *model.Employee) error {
	return r.db.Create(employee).Error
}

// Phương thức Update mới sử dụng db.Updates cho partial update
func (r *employeeRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Employee{}).Where("id = ?", id).Updates(updates).Error
}

func (r *employeeRepo) Delete(id uint) error {
	return r.db.Delete(&model.Employee{}, id).Error
}
