package service

import (
	"manage_restaurent/internal/dto"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type RoleService struct {
	repo *repository.RoleRepo
}

func NewRoleService(r *repository.RoleRepo) *RoleService {
	return &RoleService{repo: r}
}

func (s *RoleService) Create(roleDTO *dto.CreateRoleDTO) (*model.Role, error) {
	role := &model.Role{
		RoleName: roleDTO.RoleName,
	}

	if len(roleDTO.PermissionIDs) > 0 {
		var permissions []model.Permission
		if err := s.repo.GetDB().Where("id IN ?", roleDTO.PermissionIDs).Find(&permissions).Error; err != nil {
			return nil, err
		}
		role.Permissions = &permissions
	}

	if err := s.repo.Create(role); err != nil {
		return nil, err
	}

	return role, nil
}

func (s *RoleService) GetByID(id uint) (*model.Role, error) {
	return s.repo.GetByID(id)
}

func (s *RoleService) Update(id uint, updateDTO *dto.UpdateRoleDTO) (*model.Role, error) {
	// Start transaction
	tx := s.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get role with current permissions
	var role model.Role
	if err := tx.Preload("Permissions").First(&role, id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update role name if provided
	if updateDTO.RoleName != "" {
		role.RoleName = updateDTO.RoleName
		if err := tx.Save(&role).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Update permissions if provided
	if updateDTO.PermissionIDs != nil {
		// Clear existing permissions
		if err := tx.Model(&role).Association("Permissions").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		if len(updateDTO.PermissionIDs) > 0 {
			// Get new permissions
			var permissions []model.Permission
			if err := tx.Where("id IN ?", updateDTO.PermissionIDs).Find(&permissions).Error; err != nil {
				tx.Rollback()
				return nil, err
			}

			// Append new permissions
			if err := tx.Model(&role).Association("Permissions").Append(&permissions); err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Get updated role with permissions
	return s.repo.GetByID(id)
}

func (s *RoleService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *RoleService) List(offset, limit int) ([]model.Role, int64, error) {
	return s.repo.List(offset, limit)
}
