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

func (h *ShiftScheduleHandler) GetAll(c *gin.Context) {
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
