package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
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
	repo *repository.AccountRepo
}

func NewAccountService(r *repository.AccountRepo) *AccountService {
	return &AccountService{repo: r}
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

func (s *AccountService) Login(username, password string) (string, error) {
	acc, err := s.repo.GetByUserName(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}
	claims := jwt.MapClaims{
		"user_id": acc.ID,
		"role_id": acc.RoleId,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}
	return tokenString, nil
} 