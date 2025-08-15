package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"manage_restaurent/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	svc *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách category
// @Description Lấy danh sách category
// @Tags category
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Category}
// @Router /categories [get]
func (h *CategoryHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	preloadFields := utils.ParsePopulateQuery(c.Request.URL.Query())

	filters := make(map[string]interface{})
	if id := c.Query("id"); id != "" {
		idInt, _ := strconv.Atoi(id)
		filters["id"] = idInt
	}
	if employeeId := c.Query("employee_id"); employeeId != "" {
		employeeIdInt, _ := strconv.Atoi(employeeId)
		filters["employee_id"] = employeeIdInt
	}
	if shiftId := c.Query("shift_id"); shiftId != "" {
		shiftIdInt, _ := strconv.Atoi(shiftId)
		filters["shift_id"] = shiftIdInt
	}

	list, total, err := h.svc.List(page, pageSize, preloadFields)
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
// @Summary Lấy chi tiết category
// @Description Lấy chi tiết một category theo ID
// @Tags category
// @Produce json
// @Param id path int true "ID category"
// @Success 200 {object} model.Category
// @Failure 404 {object} response.ErrorResponse
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	category, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Category not found")
		return
	}

	response.Success(c, category, nil)
}

// Create godoc
// @Summary Tạo mới category
// @Description Tạo mới một category
// @Tags category
// @Accept json
// @Produce json
// @Param category body model.Category true "Dữ liệu category" example({"name":"category_name"})
// @Success 200 {object} model.Category
// @Failure 400 {object} response.ErrorResponse
// @Router /categories [post]
func (h *CategoryHandler) Create(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.Create(&category); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, category, nil)
}

// Update godoc
// @Summary Cập nhật category
// @Description Cập nhật thông tin category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID category"
// @Param category body model.Category true "Dữ liệu cập nhật" example({"name":"category_name"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /categories/{id} [patch]
func (h *CategoryHandler) Update(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	category.ID = id
	if err := h.svc.Update(&category); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, "Category updated successfully", nil)
}

// Delete godoc
// @Summary Xóa category
// @Description Xóa một category
// @Tags category
// @Produce json
// @Param id path int true "ID category"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /categories/{id} [delete]
func (h *CategoryHandler) Delete(c *gin.Context) {
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
	response.Success(c, "Category deleted successfully", nil)
}
