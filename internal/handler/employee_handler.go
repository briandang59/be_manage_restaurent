package handler

import (
	"manage_restaurent/internal/dto"
	"manage_restaurent/internal/enum"
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

// GetAll godoc
// @Summary Lấy danh sách nhân viên
// @Description Lấy danh sách nhân viên
// @Tags employee
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Employee}
// @Router /employees [get]
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

// GetByID godoc
// @Summary Lấy chi tiết nhân viên
// @Description Lấy chi tiết một nhân viên theo ID
// @Tags employee
// @Produce json
// @Param id path int true "ID nhân viên"
// @Success 200 {object} model.Employee
// @Failure 404 {object} response.ErrorResponse
// @Router /employees/{id} [get]
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

// Create godoc
// @Summary Tạo mới nhân viên
// @Description Tạo mới một nhân viên. Nếu có role_id thì tự động tạo account
// @Tags employee
// @Accept json
// @Produce json
// @Param employee body model.Employee true "Dữ liệu nhân viên" example({"full_name":"Nguyễn Văn A","gender":true,"birthday":"1990-01-01","avatar_file_id":2})
// @Success 200 {object} model.Employee
// @Failure 400 {object} response.ErrorResponse
// @Router /employees [post]
func (h *EmployeeHandler) Create(c *gin.Context) {
	// Đọc request body một lần duy nhất
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Kiểm tra có role_id không
	if _, exists := requestBody["role_id"]; exists {
		// Có role_id -> tạo employee với account tự động
		// Chuyển đổi requestBody thành CreateEmployeeDTO
		createDTO := dto.CreateEmployeeDTO{
			FullName:      getString(requestBody, "full_name"),
			Gender:        getBool(requestBody, "gender"),
			Birthday:      getString(requestBody, "birthday"),
			PhoneNumber:   getString(requestBody, "phone_number"),
			Email:         getString(requestBody, "email"),
			ScheduleType:  getScheduleType(requestBody, "schedule_type"),
			Address:       getString(requestBody, "address"),
			JoinDate:      getString(requestBody, "join_date"),
			BaseSalary:    getInt64(requestBody, "base_salary"),
			SalaryPerHour: getInt64(requestBody, "salary_per_hour"),
			AvatarFileID:  getUintPtr(requestBody, "avatar_file_id"),
			RoleId:        getUint(requestBody, "role_id"),
		}
		
		employee, err := h.svc.CreateWithAutoAccount(&createDTO)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(c, employee, nil)
	} else {
		// Không có role_id -> tạo employee thông thường
		// Chuyển đổi requestBody thành Employee model
		employee := model.Employee{
			FullName:      getString(requestBody, "full_name"),
			Gender:        getBool(requestBody, "gender"),
			Birthday:      getString(requestBody, "birthday"),
			PhoneNumber:   getString(requestBody, "phone_number"),
			Email:         getString(requestBody, "email"),
			ScheduleType:  getScheduleType(requestBody, "schedule_type"),
			Address:       getString(requestBody, "address"),
			JoinDate:      getString(requestBody, "join_date"),
			BaseSalary:    getInt64(requestBody, "base_salary"),
			SalaryPerHour: getInt64(requestBody, "salary_per_hour"),
			AvatarFileID:  getUintPtr(requestBody, "avatar_file_id"),
		}
		
		if err := h.svc.Create(&employee); err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(c, employee, nil)
	}
}

// Helper functions để chuyển đổi kiểu dữ liệu
func getString(data map[string]interface{}, key string) string {
	if val, exists := data[key]; exists {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getBool(data map[string]interface{}, key string) bool {
	if val, exists := data[key]; exists {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}

func getInt64(data map[string]interface{}, key string) int64 {
	if val, exists := data[key]; exists {
		switch v := val.(type) {
		case float64:
			return int64(v)
		case int:
			return int64(v)
		case int64:
			return v
		}
	}
	return 0
}

func getUint(data map[string]interface{}, key string) uint {
	if val, exists := data[key]; exists {
		switch v := val.(type) {
		case float64:
			return uint(v)
		case int:
			return uint(v)
		case uint:
			return v
		}
	}
	return 0
}

func getUintPtr(data map[string]interface{}, key string) *uint {
	if val, exists := data[key]; exists && val != nil {
		var u uint
		switch v := val.(type) {
		case float64:
			u = uint(v)
		case int:
			u = uint(v)
		case uint:
			u = v
		default:
			return nil
		}
		return &u
	}
	return nil
}

func getScheduleType(data map[string]interface{}, key string) enum.EmployeeScheduleType {
	if val, exists := data[key]; exists {
		if str, ok := val.(string); ok {
			return enum.EmployeeScheduleType(str)
		}
	}
	return ""
}

// CreateWithAutoAccount godoc
// @Summary Tạo mới nhân viên với account tự động
// @Description Tạo mới một nhân viên và tự động tạo account với username theo định dạng yymmdd và password mặc định
// @Tags employee
// @Accept json
// @Produce json
// @Param employee body dto.CreateEmployeeDTO true "Dữ liệu nhân viên với role_id" example({"full_name":"Nguyễn Văn A","gender":true,"birthday":"1990-01-01","role_id":1})
// @Success 200 {object} model.Employee
// @Failure 400 {object} response.ErrorResponse
// @Router /employees/with-account [post]
func (h *EmployeeHandler) CreateWithAutoAccount(c *gin.Context) {
	var createDTO dto.CreateEmployeeDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	
	employee, err := h.svc.CreateWithAutoAccount(&createDTO)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, employee, nil)
}

// Update godoc
// @Summary Cập nhật nhân viên
// @Description Cập nhật thông tin một nhân viên. Nếu có role_id thì cập nhật account
// @Tags employee
// @Accept json
// @Produce json
// @Param id path int true "ID nhân viên"
// @Param updates body object true "Dữ liệu cập nhật" example({"full_name":"Nguyễn Văn B","role_id":2})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /employees/{id} [patch]
func (h *EmployeeHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Kiểm tra có role_id không
	if _, exists := requestBody["role_id"]; exists {
		// Có role_id -> cập nhật employee và account
		if err := h.svc.UpdateWithAccount(uint(id), requestBody); err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		// Không có role_id -> cập nhật employee thông thường
		if err := h.svc.Update(uint(id), requestBody); err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	
	response.Success(c, "Employee updated successfully", nil)
}

// Delete godoc
// @Summary Xóa nhân viên
// @Description Xóa một nhân viên
// @Tags employee
// @Produce json
// @Param id path int true "ID nhân viên"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /employees/{id} [delete]
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
