package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type MenuItemRepo struct {
	db *gorm.DB
}

func NewMenuItemRepo(db *gorm.DB) *MenuItemRepo {
	return &MenuItemRepo{db: db}
}

func (r *MenuItemRepo) Create(menuItem *model.MenuItem) error {
	return r.db.Create(menuItem).Error
}

func (r *MenuItemRepo) BulkCreate(menuItems []model.MenuItem) error {
	return r.db.Create(&menuItems).Error
}

func (r *MenuItemRepo) GetByID(id uint) (*model.MenuItem, error) {
	var menuItem model.MenuItem
	if err := r.db.First(&menuItem, id).Error; err != nil {
		return nil, err
	}
	return &menuItem, nil
}

func (r *MenuItemRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.MenuItem{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MenuItemRepo) Delete(id uint) error {
	return r.db.Delete(&model.MenuItem{}, id).Error
}

func (r *MenuItemRepo) List(offset, limit int, preloadFields []string) ([]model.MenuItem, int64, error) {
	var menuItems []model.MenuItem
	var total int64
	
	query := r.db.Model(&model.MenuItem{})
	for _, field := range preloadFields {
		query = query.Preload(field)
	}
	
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	if err := query.Offset(offset).Limit(limit).Find(&menuItems).Error; err != nil {
		return nil, 0, err
	}
	return menuItems, total, nil
} 