package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuItemHandler struct {
	svc *service.MenuItemService
}

func NewMenuItemHandler(s *service.MenuItemService) *MenuItemHandler {
	return &MenuItemHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách menu item
// @Description Lấy danh sách các món ăn trong thực đơn
// @Tags menuitem
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.MenuItem}
// @Router /menu-items [get]
func (h *MenuItemHandler) GetAll(c *gin.Context) {
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
// @Summary Lấy chi tiết menu item
// @Description Lấy chi tiết một món ăn theo ID
// @Tags menuitem
// @Produce json
// @Param id path int true "ID menu item"
// @Success 200 {object} model.MenuItem
// @Failure 404 {object} response.ErrorResponse
// @Router /menu-items/{id} [get]
func (h *MenuItemHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	menuItem, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "MenuItem not found")
		return
	}
	response.Success(c, menuItem, nil)
}

// Create godoc
// @Summary Tạo mới menu item
// @Description Tạo mới một món ăn trong thực đơn
// @Tags menuitem
// @Accept json
// @Produce json
// @Param menuItem body model.MenuItem true "Dữ liệu menu item" example({"name":"Pizza Margherita","description":"Pizza truyền thống Ý","price":120000,"file_id":1,"status":"Available"})
// @Success 200 {object} model.MenuItem
// @Failure 400 {object} response.ErrorResponse
// @Router /menu-items [post]
func (h *MenuItemHandler) Create(c *gin.Context) {
	var menuItem model.MenuItem
	if err := c.ShouldBindJSON(&menuItem); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&menuItem); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, menuItem, nil)
}

// Update godoc
// @Summary Cập nhật menu item
// @Description Cập nhật thông tin một món ăn
// @Tags menuitem
// @Accept json
// @Produce json
// @Param id path int true "ID menu item"
// @Param updates body object true "Dữ liệu cập nhật" example({"name":"Pizza Hải sản","price":150000})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /menu-items/{id} [patch]
func (h *MenuItemHandler) Update(c *gin.Context) {
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
	response.Success(c, "MenuItem updated successfully", nil)
}

// Delete godoc
// @Summary Xóa menu item
// @Description Xóa một món ăn khỏi thực đơn
// @Tags menuitem
// @Produce json
// @Param id path int true "ID menu item"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /menu-items/{id} [delete]
func (h *MenuItemHandler) Delete(c *gin.Context) {
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
	response.Success(c, "MenuItem deleted successfully", nil)
} 