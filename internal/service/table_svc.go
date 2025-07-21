package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type TableService struct {
	repo repository.TableRepo
}

func NewTableService(r repository.TableRepo) *TableService {
	return &TableService{repo: r}
}
func (s *TableService) GetAll(page, pageSize int, preloadFields []string) ([]model.Table, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields)
}

func (s *TableService) Create(availibility *model.Table) error {
	return s.repo.Create(availibility)
}

func (s *TableService) Update(id uint, updates map[string]interface{}) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Update(id, updates)
}

func (s *TableService) Delete(id uint) error {
	return s.repo.Delete(id)
}
