package service

import (
	"fmt"
	"manage_restaurent/internal/dto"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
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

// CreateWithAutoAccount tạo employee cùng với account tự động
func (s *EmployeeService) CreateWithAutoAccount(dto *dto.CreateEmployeeDTO) (*model.Employee, error) {
	// Bắt đầu transaction
	tx := s.repo.GetDB().Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Tạo username theo định dạng yymmdd + số thứ tự
	now := time.Now()
	baseUsername := fmt.Sprintf("%02d%02d%02d", now.Year()%100, now.Month(), now.Day())

	// Tìm số thứ tự tiếp theo cho username
	var count int64
	tx.Model(&model.Account{}).Where("user_name LIKE ?", baseUsername+"%").Count(&count)
	username := fmt.Sprintf("%s%02d", baseUsername, count+1)

	// Tạo account với password cố định "123456"
	hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	account := &model.Account{
		UserName: username,
		Password: string(hash), // bcrypt hash của "123456"
		RoleId:   dto.RoleId,
	}

	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Tạo employee với account_id
	employee := &model.Employee{
		FullName:      dto.FullName,
		Gender:        dto.Gender,
		Birthday:      dto.Birthday,
		PhoneNumber:   dto.PhoneNumber,
		Email:         dto.Email,
		ScheduleType:  dto.ScheduleType,
		Address:       dto.Address,
		JoinDate:      dto.JoinDate,
		BaseSalary:    dto.BaseSalary,
		SalaryPerHour: dto.SalaryPerHour,
		AccountID:     &account.ID,
		AvatarFileID:  dto.AvatarFileID,
	}

	if err := tx.Create(employee).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return employee, nil
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

// UpdateWithAccount cập nhật employee và account
func (s *EmployeeService) UpdateWithAccount(id uint, updates map[string]interface{}) error {
	// Bắt đầu transaction
	tx := s.repo.GetDB().Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Lấy employee hiện tại
	employee, err := s.repo.FindByID(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Tách role_id ra khỏi updates để cập nhật account
	roleID, hasRoleID := updates["role_id"]
	delete(updates, "role_id") // Xóa role_id khỏi updates để không cập nhật vào employee

	// Cập nhật employee
	if err := tx.Model(&model.Employee{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Nếu có role_id và employee có account_id
	if hasRoleID && employee.AccountID != nil {
		// Cập nhật account
		if err := tx.Model(&model.Account{}).Where("id = ?", *employee.AccountID).Update("role_id", roleID).Error; err != nil {
			tx.Rollback()
			return err
		}
		fmt.Printf("Updated account ID %d with role_id %v\n", *employee.AccountID, roleID)
	} else if hasRoleID && employee.AccountID == nil {
		// Nếu employee chưa có account, tạo mới account
		now := time.Now()
		baseUsername := fmt.Sprintf("%02d%02d%02d", now.Year()%100, now.Month(), now.Day())

		// Tìm số thứ tự tiếp theo cho username
		var count int64
		tx.Model(&model.Account{}).Where("user_name LIKE ?", baseUsername+"%").Count(&count)
		username := fmt.Sprintf("%s%02d", baseUsername, count+1)

		hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		account := &model.Account{
			UserName: username,
			Password: string(hash),
			RoleId:   getUintFromInterface(roleID),
		}

		if err := tx.Create(account).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Cập nhật employee với account_id mới
		if err := tx.Model(&model.Employee{}).Where("id = ?", id).Update("account_id", account.ID).Error; err != nil {
			tx.Rollback()
			return err
		}

		fmt.Printf("Created new account ID %d for employee ID %d\n", account.ID, id)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// Helper function để chuyển đổi interface{} thành uint
func getUintFromInterface(val interface{}) uint {
	switch v := val.(type) {
	case float64:
		return uint(v)
	case int:
		return uint(v)
	case uint:
		return v
	default:
		return 0
	}
}

func (s *EmployeeService) Delete(id uint) error {
	return s.repo.Delete(id)
}
