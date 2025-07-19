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

// AvailibilityHandler xử lý các request HTTP cho Availibility
type AvailibilityHandler struct {
	svc *service.AvailibilityService
}

// NewAvailibilityHandler tạo một thể hiện mới của AvailibilityHandler
func NewAvailibilityHandler(s *service.AvailibilityService) *AvailibilityHandler {
	return &AvailibilityHandler{svc: s}
}

func (h *AvailibilityHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	preloadFields := utils.ParsePopulateQuery(c.Request.URL.Query())

	list, total, err := h.svc.GetAll(page, pageSize, preloadFields)
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

func (h *AvailibilityHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	availibility, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Availibility not found")
		return
	}

	response.Success(c, availibility, nil)
}

func (h *AvailibilityHandler) Create(c *gin.Context) {
	// Hỗ trợ tạo một hoặc nhiều bản ghi
	var availabilities []model.Availibility
	if err := c.ShouldBindJSON(&availabilities); err != nil {
		var availibility model.Availibility
		if err := c.ShouldBindJSON(&availibility); err != nil {
			response.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		availabilities = append(availabilities, availibility)
	}

	if err := h.svc.BulkCreate(availabilities); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, availabilities, nil)
}

func (h *AvailibilityHandler) Update(c *gin.Context) {
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
	response.Success(c, "Availibility updated successfully", nil)
}

func (h *AvailibilityHandler) Delete(c *gin.Context) {
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
	response.Success(c, "Availibility deleted successfully", nil)
}
