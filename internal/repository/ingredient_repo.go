package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type IngredientRepo struct {
	db *gorm.DB
}

func NewIngredientRepo(db *gorm.DB) *IngredientRepo {
	return &IngredientRepo{db: db}
}

func (r *IngredientRepo) Create(ingredient *model.Ingredient) error {
	return r.db.Create(ingredient).Error
}

func (r *IngredientRepo) GetByID(id uint) (*model.Ingredient, error) {
	var ingredient model.Ingredient
	if err := r.db.First(&ingredient, id).Error; err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *IngredientRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Ingredient{}).Where("id = ?", id).Updates(updates).Error
}

func (r *IngredientRepo) Delete(id uint) error {
	return r.db.Delete(&model.Ingredient{}, id).Error
}

func (r *IngredientRepo) List(offset, limit int) ([]model.Ingredient, int64, error) {
	var ingredients []model.Ingredient
	var total int64
	r.db.Model(&model.Ingredient{}).Count(&total)
	if err := r.db.Offset(offset).Limit(limit).Find(&ingredients).Error; err != nil {
		return nil, 0, err
	}
	return ingredients, total, nil
} 