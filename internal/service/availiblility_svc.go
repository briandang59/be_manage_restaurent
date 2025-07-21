package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type AvailibilityService struct {
	repo repository.AvailibilityRepo
}

func NewAvailibilityService(r repository.AvailibilityRepo) *AvailibilityService {
	return &AvailibilityService{repo: r}
}

func (s *AvailibilityService) GetAll(page, pageSize int, preloadFields []string, filters map[string]interface{}) ([]model.Availibility, int64, error) {
	return s.repo.FindAll(page, pageSize, preloadFields, filters)
}
func (s *AvailibilityService) GetByID(id uint) (*model.Availibility, error) {
	return s.repo.FindByID(id)
}

func (s *AvailibilityService) Create(availibility *model.Availibility) error {
	return s.repo.Create(availibility)
}

func (s *AvailibilityService) BulkCreate(availibilities []model.Availibility) error {
	return s.repo.BulkCreate(availibilities)
}

func (s *AvailibilityService) Update(id uint, updates map[string]interface{}) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Update(id, updates)
}

func (s *AvailibilityService) Delete(id uint) error {
	return s.repo.Delete(id)
}
