package service

import (
	"errors"
	"os"
	"time"

	"manage_restaurent/internal/dto"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// getJWTSecret lấy JWT secret từ biến môi trường
func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Fallback secret nếu không có biến môi trường (chỉ dùng cho development)
		secret = "your_secret_key_development_only"
	}
	return []byte(secret)
}

type AccountService struct {
	repo         *repository.AccountRepo
	employeeRepo repository.EmployeeRepo
}

func NewAccountService(r *repository.AccountRepo, employeeRepo repository.EmployeeRepo) *AccountService {
	return &AccountService{repo: r, employeeRepo: employeeRepo}
}

func (s *AccountService) Create(account *model.Account) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	account.Password = string(hash)
	return s.repo.Create(account)
}

func (s *AccountService) GetByID(id uint) (*model.Account, error) {
	return s.repo.GetByID(id)
}

func (s *AccountService) GetByUserName(username string) (*model.Account, error) {
	return s.repo.GetByUserName(username)
}

func (s *AccountService) Update(id uint, updates map[string]interface{}) error {
	if pwd, ok := updates["password"]; ok {
		hash, err := bcrypt.GenerateFromPassword([]byte(pwd.(string)), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updates["password"] = string(hash)
	}
	return s.repo.Update(id, updates)
}

func (s *AccountService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AccountService) List(offset, limit int, preloadFields []string) ([]model.Account, int64, error) {
	return s.repo.List(offset, limit, preloadFields)
}

func (s *AccountService) Login(username, password string) (*dto.LoginResponseDTO, error) {
	acc, err := s.repo.GetByUserNameWithRole(username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Kiểm tra password với bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}
	claims := jwt.MapClaims{
		"user_id": acc.ID,
		"role_id": acc.RoleId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return nil, err
	}

	// Create user info without password
	userInfo := map[string]interface{}{
		"id":         acc.ID,
		"user_name":  acc.UserName,
		"role_id":    acc.RoleId,
		"created_at": acc.CreatedAt,
		"updated_at": acc.UpdatedAt,
	}

	// Create role info
	var roleInfo interface{}
	if acc.Role != nil {
		roleInfo = map[string]interface{}{
			"id":        acc.Role.ID,
			"role_name": acc.Role.RoleName,
		}
	}

	// Get employee info if exists
	var employeeInfo interface{}
	employee, err := s.employeeRepo.FindByAccountID(acc.ID)
	if err == nil && employee != nil {
		employeeInfo = map[string]interface{}{
			"id":              employee.ID,
			"full_name":       employee.FullName,
			"gender":          employee.Gender,
			"birthday":        employee.Birthday,
			"phone_number":    employee.PhoneNumber,
			"email":           employee.Email,
			"schedule_type":   employee.ScheduleType,
			"address":         employee.Address,
			"join_date":       employee.JoinDate,
			"base_salary":     employee.BaseSalary,
			"salary_per_hour": employee.SalaryPerHour,
			"avatar_file_id":  employee.AvatarFileID,
		}
	}

	return &dto.LoginResponseDTO{
		Token:    tokenString,
		User:     userInfo,
		Role:     roleInfo,
		Employee: employeeInfo,
	}, nil
}
