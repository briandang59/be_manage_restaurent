package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type MenuItemService struct {
	repo *repository.MenuItemRepo
}

func NewMenuItemService(r *repository.MenuItemRepo) *MenuItemService {
	return &MenuItemService{repo: r}
}

func (s *MenuItemService) Create(menuItem *model.MenuItem) error {
	return s.repo.Create(menuItem)
}

func (s *MenuItemService) GetByID(id uint) (*model.MenuItem, error) {
	return s.repo.GetByID(id)
}

func (s *MenuItemService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *MenuItemService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *MenuItemService) List(offset, limit int, preloadFields []string) ([]model.MenuItem, int64, error) {
	return s.repo.List(offset, limit, preloadFields)
}

func (s *MenuItemService) BulkCreate(menuItems []model.MenuItem) error {
	return s.repo.BulkCreate(menuItems)
} 