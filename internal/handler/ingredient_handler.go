package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IngredientHandler struct {
	svc *service.IngredientService
}

func NewIngredientHandler(s *service.IngredientService) *IngredientHandler {
	return &IngredientHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách nguyên liệu
// @Description Lấy danh sách nguyên liệu
// @Tags ingredient
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Ingredient}
// @Router /ingredients [get]
func (h *IngredientHandler) GetAll(c *gin.Context) {
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
// @Summary Lấy chi tiết nguyên liệu
// @Description Lấy chi tiết một nguyên liệu theo ID
// @Tags ingredient
// @Produce json
// @Param id path int true "ID nguyên liệu"
// @Success 200 {object} model.Ingredient
// @Failure 404 {object} response.ErrorResponse
// @Router /ingredients/{id} [get]
func (h *IngredientHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	ingredient, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Ingredient not found")
		return
	}
	response.Success(c, ingredient, nil)
}

// Create godoc
// @Summary Tạo mới nguyên liệu
// @Description Tạo mới một nguyên liệu
// @Tags ingredient
// @Accept json
// @Produce json
// @Param ingredient body model.Ingredient true "Dữ liệu nguyên liệu" example({"name":"Bột mì","quantity":100,"unit":"kg","warning_quantity":10})
// @Success 200 {object} model.Ingredient
// @Failure 400 {object} response.ErrorResponse
// @Router /ingredients [post]
func (h *IngredientHandler) Create(c *gin.Context) {
	var ingredient model.Ingredient
	if err := c.ShouldBindJSON(&ingredient); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&ingredient); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, ingredient, nil)
}

// Update godoc
// @Summary Cập nhật nguyên liệu
// @Description Cập nhật thông tin một nguyên liệu
// @Tags ingredient
// @Accept json
// @Produce json
// @Param id path int true "ID nguyên liệu"
// @Param updates body object true "Dữ liệu cập nhật" example({"quantity":200})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /ingredients/{id} [patch]
func (h *IngredientHandler) Update(c *gin.Context) {
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
	response.Success(c, "Ingredient updated successfully", nil)
}

// Delete godoc
// @Summary Xóa nguyên liệu
// @Description Xóa một nguyên liệu
// @Tags ingredient
// @Produce json
// @Param id path int true "ID nguyên liệu"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /ingredients/{id} [delete]
func (h *IngredientHandler) Delete(c *gin.Context) {
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
	response.Success(c, "Ingredient deleted successfully", nil)
} 