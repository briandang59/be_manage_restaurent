package dto

type CreateApplyRecruitmentDTO struct {
	RecruitmentID uint   `json:"recruitment_id" binding:"required"`
	FullName      string `json:"fullname"     binding:"required"`
	Email         string `json:"email"        binding:"required,email"`
	PhoneNumber   string `json:"phone_number" binding:"required"`
	CVID          uint   `json:"cv_id"        binding:"required"`
}

type UpdateApplyRecruitmentDTO struct {
	FullName    *string `json:"fullname"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	CVID        *uint   `json:"cv_id"`
}
