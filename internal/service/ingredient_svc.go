package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type IngredientService struct {
	repo *repository.IngredientRepo
}

func NewIngredientService(r *repository.IngredientRepo) *IngredientService {
	return &IngredientService{repo: r}
}

func (s *IngredientService) Create(ingredient *model.Ingredient) error {
	return s.repo.Create(ingredient)
}

func (s *IngredientService) GetByID(id uint) (*model.Ingredient, error) {
	return s.repo.GetByID(id)
}

func (s *IngredientService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *IngredientService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *IngredientService) List(offset, limit int) ([]model.Ingredient, int64, error) {
	return s.repo.List(offset, limit)
} 