package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type RoleService struct {
	repo *repository.RoleRepo
}

func NewRoleService(r *repository.RoleRepo) *RoleService {
	return &RoleService{repo: r}
}

func (s *RoleService) Create(role *model.Role) error {
	return s.repo.Create(role)
}

func (s *RoleService) GetByID(id uint) (*model.Role, error) {
	return s.repo.GetByID(id)
}

func (s *RoleService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *RoleService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *RoleService) List(offset, limit int) ([]model.Role, int64, error) {
	return s.repo.List(offset, limit)
} 