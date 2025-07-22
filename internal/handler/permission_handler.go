package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	svc *service.PermissionService
}

func NewPermissionHandler(s *service.PermissionService) *PermissionHandler {
	return &PermissionHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách permission
// @Description Lấy danh sách quyền
// @Tags permission
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Permission}
// @Router /permissions [get]
func (h *PermissionHandler) GetAll(c *gin.Context) {
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
// @Summary Lấy chi tiết permission
// @Description Lấy chi tiết một quyền theo ID
// @Tags permission
// @Produce json
// @Param id path int true "ID permission"
// @Success 200 {object} model.Permission
// @Failure 404 {object} response.ErrorResponse
// @Router /permissions/{id} [get]
func (h *PermissionHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	permission, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Permission not found")
		return
	}
	response.Success(c, permission, nil)
}

// Create godoc
// @Summary Tạo mới permission
// @Description Tạo mới một quyền
// @Tags permission
// @Accept json
// @Produce json
// @Param permission body model.Permission true "Dữ liệu permission" example({"permission_name":"view_menu"})
// @Success 200 {object} model.Permission
// @Failure 400 {object} response.ErrorResponse
// @Router /permissions [post]
func (h *PermissionHandler) Create(c *gin.Context) {
	var permission model.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&permission); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, permission, nil)
}

// Update godoc
// @Summary Cập nhật permission
// @Description Cập nhật thông tin một quyền
// @Tags permission
// @Accept json
// @Produce json
// @Param id path int true "ID permission"
// @Param updates body object true "Dữ liệu cập nhật" example({"permission_name":"edit_menu"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /permissions/{id} [patch]
func (h *PermissionHandler) Update(c *gin.Context) {
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
	response.Success(c, "Permission updated successfully", nil)
}

// Delete godoc
// @Summary Xóa permission
// @Description Xóa một quyền
// @Tags permission
// @Produce json
// @Param id path int true "ID permission"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /permissions/{id} [delete]
func (h *PermissionHandler) Delete(c *gin.Context) {
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
	response.Success(c, "Permission deleted successfully", nil)
} 