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