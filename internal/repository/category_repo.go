package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	FindAll(page, pageSize int, preloadFields []string) ([]model.Category, int64, error)
	Create(category *model.Category) error
	FindByID(id uint) (*model.Category, error)
	Update(category *model.Category) error
	Delete(id uint) error
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.Category, int64, error) {
	var list []model.Category
	var total int64
	offset := (page - 1) * pageSize

	query := r.db.Model(&model.Category{})

	for _, field := range preloadFields {
		query = query.Preload(field)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *categoryRepo) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepo) FindByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepo) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepo) Delete(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}
