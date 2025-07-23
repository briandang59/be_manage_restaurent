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

// ShiftScheduleHandler xử lý các request HTTP cho ShiftSchedule
type ShiftScheduleHandler struct {
	svc *service.ShiftScheduleService
}

// NewShiftScheduleHandler tạo một thể hiện mới của ShiftScheduleHandler
func NewShiftScheduleHandler(s *service.ShiftScheduleService) *ShiftScheduleHandler {
	return &ShiftScheduleHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách lịch ca làm việc
// @Description Lấy danh sách lịch ca làm việc
// @Tags shiftschedule
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.ShiftSchedule}
// @Router /shifts-chedules [get]
func (h *ShiftScheduleHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	filters := make(map[string]interface{})

	if employeeId := c.Query("employee_id"); employeeId != "" {
		employeeIdInt, _ := strconv.Atoi(employeeId)
		filters["employee_id"] = employeeIdInt
	}
	if shiftId := c.Query("shift_id"); shiftId != "" {
		shiftIdInt, _ := strconv.Atoi(shiftId)
		filters["shift_id"] = shiftIdInt
	}

	preloadFields := utils.ParsePopulateQuery(c.Request.URL.Query())

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
// @Summary Lấy chi tiết lịch ca làm việc
// @Description Lấy chi tiết một lịch ca làm việc theo ID
// @Tags shiftschedule
// @Produce json
// @Param id path int true "ID shift schedule"
// @Success 200 {object} model.ShiftSchedule
// @Failure 404 {object} response.ErrorResponse
// @Router /shifts-chedules/{id} [get]
func (h *ShiftScheduleHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	shiftSchedule, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "ShiftSchedule not found")
		return
	}

	response.Success(c, shiftSchedule, nil)
}

// Create godoc
// @Summary Tạo mới lịch ca làm việc
// @Description Tạo mới một hoặc nhiều lịch ca làm việc
// @Tags shiftschedule
// @Accept json
// @Produce json
// @Param shiftSchedules body []model.ShiftSchedule true "Dữ liệu lịch ca làm việc" example([{ "employee_id": 1, "shift_id": 2, "date": "2024-07-22" }])
// @Success 200 {object} []model.ShiftSchedule
// @Failure 400 {object} response.ErrorResponse
// @Router /shifts-chedules [post]
func (h *ShiftScheduleHandler) Create(c *gin.Context) {
	var shiftSchedules []model.ShiftSchedule
	if err := c.ShouldBindJSON(&shiftSchedules); err != nil {
		var shiftSchedule model.ShiftSchedule
		if err := c.ShouldBindJSON(&shiftSchedule); err != nil {
			response.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		shiftSchedules = append(shiftSchedules, shiftSchedule)
	}

	if err := h.svc.BulkCreate(shiftSchedules); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, shiftSchedules, nil)
}

// Update godoc
// @Summary Cập nhật lịch ca làm việc
// @Description Cập nhật thông tin lịch ca làm việc
// @Tags shiftschedule
// @Accept json
// @Produce json
// @Param id path int true "ID shift schedule"
// @Param updates body object true "Dữ liệu cập nhật" example({"date":"2024-07-23"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /shifts-chedules/{id} [patch]
func (h *ShiftScheduleHandler) Update(c *gin.Context) {
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
	response.Success(c, "ShiftSchedule updated successfully", nil)
}

// Delete godoc
// @Summary Xóa lịch ca làm việc
// @Description Xóa một lịch ca làm việc
// @Tags shiftschedule
// @Produce json
// @Param id path int true "ID shift schedule"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /shifts-chedules/{id} [delete]
func (h *ShiftScheduleHandler) Delete(c *gin.Context) {
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
	response.Success(c, "ShiftSchedule deleted successfully", nil)
}
