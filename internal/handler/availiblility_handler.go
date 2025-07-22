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

// GetAll godoc
// @Summary Lấy danh sách availibility
// @Description Lấy danh sách availibility
// @Tags availibility
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Availibility}
// @Router /availibilities [get]
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

	list, total, err := h.svc.GetAll(page, pageSize, preloadFields, filters)
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
// @Summary Lấy chi tiết availibility
// @Description Lấy chi tiết một availibility theo ID
// @Tags availibility
// @Produce json
// @Param id path int true "ID availibility"
// @Success 200 {object} model.Availibility
// @Failure 404 {object} response.ErrorResponse
// @Router /availibilities/{id} [get]
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

// Create godoc
// @Summary Tạo mới availibility
// @Description Tạo mới một availibility
// @Tags availibility
// @Accept json
// @Produce json
// @Param availibility body model.Availibility true "Dữ liệu availibility" example({"employee_id":1,"shift_id":2,"date":"2024-07-22"})
// @Success 200 {object} model.Availibility
// @Failure 400 {object} response.ErrorResponse
// @Router /availibilities [post]
func (h *AvailibilityHandler) Create(c *gin.Context) {
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

// Update godoc
// @Summary Cập nhật availibility
// @Description Cập nhật thông tin availibility
// @Tags availibility
// @Accept json
// @Produce json
// @Param id path int true "ID availibility"
// @Param updates body object true "Dữ liệu cập nhật" example({"date":"2024-07-23"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /availibilities/{id} [patch]
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

// Delete godoc
// @Summary Xóa availibility
// @Description Xóa một availibility
// @Tags availibility
// @Produce json
// @Param id path int true "ID availibility"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /availibilities/{id} [delete]
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
