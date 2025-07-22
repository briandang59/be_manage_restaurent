package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	svc *service.RoleService
}

func NewRoleHandler(s *service.RoleService) *RoleHandler {
	return &RoleHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách role
// @Description Lấy danh sách quyền
// @Tags role
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Role}
// @Router /roles [get]
func (h *RoleHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	list, total, err := h.svc.List(offset, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list, &response.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
	})
}

// GetByID godoc
// @Summary Lấy chi tiết role
// @Description Lấy chi tiết một quyền theo ID
// @Tags role
// @Produce json
// @Param id path int true "ID role"
// @Success 200 {object} model.Role
// @Failure 404 {object} response.ErrorResponse
// @Router /roles/{id} [get]
func (h *RoleHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	role, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Role not found")
		return
	}
	response.Success(c, role, nil)
}

// Create godoc
// @Summary Tạo mới role
// @Description Tạo mới một quyền
// @Tags role
// @Accept json
// @Produce json
// @Param role body model.Role true "Dữ liệu role" example({"role_name":"admin"})
// @Success 200 {object} model.Role
// @Failure 400 {object} response.ErrorResponse
// @Router /roles [post]
func (h *RoleHandler) Create(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&role); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, role, nil)
}

// Update godoc
// @Summary Cập nhật role
// @Description Cập nhật thông tin một quyền
// @Tags role
// @Accept json
// @Produce json
// @Param id path int true "ID role"
// @Param updates body object true "Dữ liệu cập nhật" example({"role_name":"manager"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /roles/{id} [patch]
func (h *RoleHandler) Update(c *gin.Context) {
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
	response.Success(c, "Role updated successfully", nil)
}

// Delete godoc
// @Summary Xóa role
// @Description Xóa một quyền
// @Tags role
// @Produce json
// @Param id path int true "ID role"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /roles/{id} [delete]
func (h *RoleHandler) Delete(c *gin.Context) {
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
	response.Success(c, "Role deleted successfully", nil)
} 