package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type TableService struct {
	repo      repository.TableRepo
	orderRepo *repository.OrderRepo // Inject để check occupied status
}

func NewTableService(r repository.TableRepo, orderRepo *repository.OrderRepo) *TableService {
	return &TableService{
		repo:      r,
		orderRepo: orderRepo,
	}
}

func (s *TableService) GetAll(page, pageSize int, preloadFields []string) ([]model.Table, int64, error) {
	list, total, err := s.repo.FindAll(page, pageSize, preloadFields)
	if err != nil {
		return nil, 0, err
	}

	// Enrich each table with occupied status (true nếu có order UnPaid)
	for i := range list {
		_, err := s.orderRepo.FindOrderByTableId(list[i].ID)
		if err == nil {
			list[i].IsOccupied = true
		} else {
			list[i].IsOccupied = false
		}
	}

	return list, total, nil
}

func (s *TableService) Create(table *model.Table) error { // Sửa param name từ availibility thành table
	return s.repo.Create(table)
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
