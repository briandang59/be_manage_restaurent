package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type RecruitmentService struct {
	repo *repository.RecruitmentRepo
}

func NewRecruitmentService(r *repository.RecruitmentRepo) *RecruitmentService {
	return &RecruitmentService{repo: r}
}

func (s *RecruitmentService) Create(recuitment *model.Recruitment) error {
	return s.repo.Create(recuitment)
}
func (s *RecruitmentService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *RecruitmentService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *RecruitmentService) List(offset, limit int) ([]model.Recruitment, int64, error) {
	return s.repo.List(offset, limit)
}
