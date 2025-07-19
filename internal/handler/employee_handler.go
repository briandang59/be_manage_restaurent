package handler

import (
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"manage_restaurent/utils"
	"net/http"
	"strconv"

	"manage_restaurent/internal/model"

	"github.com/gin-gonic/gin"
)

// EmployeeHandler xử lý các request HTTP cho Employee
type EmployeeHandler struct {
	svc *service.EmployeeService
}

// NewEmployeeHandler tạo một thể hiện mới của EmployeeHandler
func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: s}
}

func (h *EmployeeHandler) GetAll(c *gin.Context) {
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

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	employee, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Employee not found")
		return
	}

	response.Success(c, employee, nil)
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&employee); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, employee, nil)
}

func (h *EmployeeHandler) Update(c *gin.Context) {
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
	response.Success(c, "Employee updated successfully", nil)
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
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
	response.Success(c, "Employee deleted successfully", nil)
}
