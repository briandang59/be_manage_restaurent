package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

func (r *RoleRepo) GetByID(id uint) (*model.Role, error) {
	var role model.Role
	if err := r.db.Preload("Permissions").First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Role{}).Where("id = ?", id).Updates(updates).Error
}

func (r *RoleRepo) Delete(id uint) error {
	return r.db.Delete(&model.Role{}, id).Error
}

func (r *RoleRepo) List(offset, limit int) ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64
	r.db.Model(&model.Role{}).Count(&total)
	if err := r.db.Preload("Permissions").Offset(offset).Limit(limit).Find(&roles).Error; err != nil {
		return nil, 0, err
	}
	return roles, total, nil
} 