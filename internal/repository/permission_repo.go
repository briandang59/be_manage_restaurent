package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type PermissionRepo struct {
	db *gorm.DB
}

func NewPermissionRepo(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{db: db}
}

func (r *PermissionRepo) Create(permission *model.Permission) error {
	return r.db.Create(permission).Error
}

func (r *PermissionRepo) GetByID(id uint) (*model.Permission, error) {
	var permission model.Permission
	if err := r.db.Preload("Roles").First(&permission, id).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Permission{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PermissionRepo) Delete(id uint) error {
	return r.db.Delete(&model.Permission{}, id).Error
}

func (r *PermissionRepo) List(offset, limit int) ([]model.Permission, int64, error) {
	var permissions []model.Permission
	var total int64
	r.db.Model(&model.Permission{}).Count(&total)
	if err := r.db.Preload("Roles").Offset(offset).Limit(limit).Find(&permissions).Error; err != nil {
		return nil, 0, err
	}
	return permissions, total, nil
} 