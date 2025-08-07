package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type PermissionService struct {
	repo *repository.PermissionRepo
}

func NewPermissionService(r *repository.PermissionRepo) *PermissionService {
	return &PermissionService{repo: r}
}

func (s *PermissionService) Create(permission *model.Permission) error {
	return s.repo.Create(permission)
}

func (s *PermissionService) GetByID(id uint) (*model.Permission, error) {
	return s.repo.GetByID(id)
}

func (s *PermissionService) Update(id uint, updates map[string]interface{}) error {
	// Convert permission_name to name in updates if it exists
	if permissionName, ok := updates["permission_name"]; ok {
		delete(updates, "permission_name")
		updates["name"] = permissionName
	}
	return s.repo.Update(id, updates)
}

func (s *PermissionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *PermissionService) List(offset, limit int) ([]model.Permission, int64, error) {
	return s.repo.List(offset, limit)
}
