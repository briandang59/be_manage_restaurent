package service

import (
	"manage_restaurent/internal/dto"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type ApplyRecruitmentService struct {
	repo *repository.ApplyRecruitmentRepo
}

func NewApplyRecruitmentService(r *repository.ApplyRecruitmentRepo) *ApplyRecruitmentService {
	return &ApplyRecruitmentService{repo: r}
}

func (s *ApplyRecruitmentService) Create(data *dto.CreateApplyRecruitmentDTO) (*model.ApplyRecruitment, error) {
	item := model.ApplyRecruitment{
		RecruitmentID: data.RecruitmentID,
		FullName:      data.FullName,
		Email:         data.Email,
		PhoneNumber:   data.PhoneNumber,
		CVID:          data.CVID,
	}

	if err := s.repo.Create(&item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *ApplyRecruitmentService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *ApplyRecruitmentService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *ApplyRecruitmentService) List(page, pageSize int, preload []string) ([]model.ApplyRecruitment, int64, error) {
	return s.repo.FindAll(page, pageSize, preload)
}
