package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepo
}

func NewCategoryService(r repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: r}
}

func (s *CategoryService) Create(category *model.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetByID(id uint) (*model.Category, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) Update(category *model.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) List(page, pageSize int, preloadFields []string) ([]model.Category, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}
