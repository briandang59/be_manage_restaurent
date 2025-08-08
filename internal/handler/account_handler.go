package handler

import (
	"manage_restaurent/internal/dto"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"manage_restaurent/utils"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	svc *service.AccountService
}

func NewAccountHandler(s *service.AccountService) *AccountHandler {
	return &AccountHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách tài khoản
// @Description Lấy danh sách tài khoản
// @Tags account
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Account}
// @Router /accounts [get]
func (h *AccountHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	preloadFields := utils.ParsePopulateQuery(c.Request.URL.Query())
	list, total, err := h.svc.List(offset, pageSize, preloadFields)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	var dtoList []dto.AccountDTO
	for _, acc := range list {
		dtoList = append(dtoList, dto.AccountDTO{
			ID:        acc.ID,
			UserName:  acc.UserName,
			RoleId:    acc.RoleId,
			Role:      acc.Role,
			CreatedAt: acc.CreatedAt,
			UpdatedAt: acc.UpdatedAt,
		})
	}
	response.Success(c, dtoList, &response.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
	})
}

// GetByID godoc
// @Summary Lấy chi tiết tài khoản
// @Description Lấy chi tiết một tài khoản theo ID
// @Tags account
// @Produce json
// @Param id path int true "ID tài khoản"
// @Success 200 {object} model.Account
// @Failure 404 {object} response.ErrorResponse
// @Router /accounts/{id} [get]
func (h *AccountHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	account, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Account not found")
		return
	}
	accDTO := dto.AccountDTO{
		ID:        account.ID,
		UserName:  account.UserName,
		RoleId:    account.RoleId,
		Role:      account.Role,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
	response.Success(c, accDTO, nil)
}

// Create godoc
// @Summary Tạo mới tài khoản
// @Description Tạo mới một tài khoản
// @Tags account
// @Accept json
// @Produce json
// @Param account body model.Account true "Dữ liệu tài khoản" example({"user_name":"admin","password":"123456","role_id":1})
// @Success 200 {object} model.Account
// @Failure 400 {object} response.ErrorResponse
// @Router /accounts [post]
func (h *AccountHandler) Create(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&account); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, account, nil)
}

// Update godoc
// @Summary Cập nhật tài khoản
// @Description Cập nhật thông tin một tài khoản
// @Tags account
// @Accept json
// @Produce json
// @Param id path int true "ID tài khoản"
// @Param updates body object true "Dữ liệu cập nhật" example({"password":"newpass"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /accounts/{id} [patch]
func (h *AccountHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Update(uint(id), updates); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, "Account updated successfully", nil)
}

// Delete godoc
// @Summary Xóa tài khoản
// @Description Xóa một tài khoản
// @Tags account
// @Produce json
// @Param id path int true "ID tài khoản"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /accounts/{id} [delete]
func (h *AccountHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, "Account deleted successfully", nil)
}

// Login godoc
// @Summary Đăng nhập
// @Description Đăng nhập tài khoản, trả về JWT token và thông tin user, role
// @Tags account
// @Accept json
// @Produce json
// @Param login body dto.LoginRequestDTO true "Thông tin đăng nhập" example({"user_name":"admin","password":"123456"})
// @Success 200 {object} dto.LoginResponseDTO
// @Failure 401 {object} response.ErrorResponse
// @Router /accounts/login [post]
func (h *AccountHandler) Login(c *gin.Context) {
	var req dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	loginResponse, err := h.svc.Login(req.UserName, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}
	response.Success(c, loginResponse, nil)
}
