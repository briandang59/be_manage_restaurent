package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type TableRepo interface {
	Create(table *model.Table) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
	FindByID(id uint) (*model.Table, error)
	FindAll(page, pageSize int, preloadFields []string) ([]model.Table, int64, error)
}

type tableRepo struct {
	db *gorm.DB
}

func NewTableRepo(db *gorm.DB) TableRepo {
	return &tableRepo{db: db}
}

func (r *tableRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Table, int64, error) {
	var list []model.Table
	var total int64
	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Table{})
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

func (r *tableRepo) Create(table *model.Table) error {
	return r.db.Create(table).Error
}

func (r *tableRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Table{}).Where("id = ?", id).Updates(updates).Error
}

func (r *tableRepo) Delete(id uint) error {
	return r.db.Delete(&model.Table{}, id).Error
}

func (r *tableRepo) FindByID(id uint) (*model.Table, error) {
	var table model.Table
	if err := r.db.First(&table, id).Error; err != nil {
		return nil, err
	}
	return &table, nil
}
