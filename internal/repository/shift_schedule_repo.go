package repository

import (
	"fmt"
	"manage_restaurent/internal/model"
	"strings"

	"gorm.io/gorm"
)

// ShiftScheduleRepo định nghĩa các phương thức tương tác với bảng ShiftSchedule
type ShiftScheduleRepo interface {
	FindAll(page, pageSize int, preloadFields []string, filters map[string]interface{}) ([]model.ShiftSchedule, int64, error)
	FindByID(id uint) (*model.ShiftSchedule, error)
	Create(shiftSchedule *model.ShiftSchedule) error
	BulkCreate(shiftSchedules []model.ShiftSchedule) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type shiftScheduleRepo struct {
	db *gorm.DB
}

// NewShiftScheduleRepo tạo một thể hiện mới của ShiftScheduleRepo
func NewShiftScheduleRepo(db *gorm.DB) ShiftScheduleRepo {
	return &shiftScheduleRepo{db: db}
}

func (r *shiftScheduleRepo) FindAll(page, pageSize int, preloadFields []string, filters map[string]interface{}) ([]model.ShiftSchedule, int64, error) {
	var list []model.ShiftSchedule
	var total int64
	offset := (page - 1) * pageSize

	// 🧹 Bước 1: Xóa các bản ghi trùng lặp (cùng employee_id, shift_id, date)
	// Giữ lại bản ghi có ID nhỏ nhất
	if err := r.db.Exec(`
		DELETE FROM shift_schedules
		WHERE id NOT IN (
			SELECT MIN(id)
			FROM shift_schedules
			GROUP BY employee_id, shift_id, date
		)
	`).Error; err != nil {
		return nil, 0, err
	}

	// 🧩 Bước 2: Xây dựng query lấy dữ liệu
	query := r.db.Model(&model.ShiftSchedule{})
	for _, field := range preloadFields {
		query = query.Preload(field)
	}

	if shiftId, ok := filters["shift_id"]; ok {
		query = query.Where("shift_id = ?", shiftId)
	}
	if employeeId, ok := filters["employee_id"]; ok {
		query = query.Where("employee_id = ?", employeeId)
	}

	// 📊 Đếm tổng số dòng
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 📥 Lấy dữ liệu có phân trang
	if err := query.
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *shiftScheduleRepo) FindByID(id uint) (*model.ShiftSchedule, error) {
	var shiftSchedule model.ShiftSchedule
	if err := r.db.First(&shiftSchedule, id).Error; err != nil {
		return nil, err
	}
	return &shiftSchedule, nil
}

func (r *shiftScheduleRepo) Create(shiftSchedule *model.ShiftSchedule) error {
	return r.db.Create(shiftSchedule).Error
}

func (r *shiftScheduleRepo) BulkCreate(shiftSchedules []model.ShiftSchedule) error {
	// 1. Thu thập tất cả các cặp EmployeeID và Date từ dữ liệu đầu vào
	var existingChecks []struct {
		EmployeeID uint
		Date       string
	}

	// Lặp qua dữ liệu đầu vào để tạo danh sách cần kiểm tra
	for _, schedule := range shiftSchedules {
		if schedule.EmployeeID != nil && schedule.Date != "" {
			existingChecks = append(existingChecks, struct {
				EmployeeID uint
				Date       string
			}{
				EmployeeID: *schedule.EmployeeID,
				Date:       schedule.Date,
			})
		}
	}

	if len(existingChecks) == 0 {
		return nil // Không có dữ liệu hợp lệ để tạo
	}

	// 2. Xây dựng truy vấn để tìm các bản ghi đã tồn tại
	// Sử dụng OR kết hợp các điều kiện (employee_id = ? AND date = ?)
	query := r.db.Model(&model.ShiftSchedule{})
	var subQueries []string
	var args []interface{}

	for _, check := range existingChecks {
		subQueries = append(subQueries, "(employee_id = ? AND date = ?)")
		args = append(args, check.EmployeeID, check.Date)
	}

	// Kết hợp tất cả các điều kiện phụ bằng OR
	whereClause := "(" + strings.Join(subQueries, " OR ") + ")"

	// 3. Truy vấn DB để lấy các bản ghi đã tồn tại
	var existingRecords []model.ShiftSchedule
	err := query.Where(whereClause, args...).Find(&existingRecords).Error
	if err != nil {
		return err // Lỗi khi truy vấn
	}

	// 4. Tạo Map các bản ghi đã tồn tại để tra cứu nhanh
	existingMap := make(map[string]bool)
	for _, record := range existingRecords {
		// Key: "employeeID_date" (ví dụ: "2_2025-10-17")
		key := fmt.Sprintf("%d_%s", *record.EmployeeID, record.Date)
		existingMap[key] = true
	}

	// 5. Lọc danh sách đầu vào, chỉ giữ lại các bản ghi CHƯA tồn tại
	var newSchedulesToCreate []model.ShiftSchedule
	for _, schedule := range shiftSchedules {
		if schedule.EmployeeID != nil && schedule.Date != "" {
			key := fmt.Sprintf("%d_%s", *schedule.EmployeeID, schedule.Date)
			if !existingMap[key] {
				newSchedulesToCreate = append(newSchedulesToCreate, schedule)
				// Đánh dấu vào map để tránh trùng lặp trong chính file đầu vào
				existingMap[key] = true
			}
		}
	}

	if len(newSchedulesToCreate) == 0 {
		// fmt.Println("Không có bản ghi mới nào để chèn.")
		return nil
	}

	// 6. Thực hiện Bulk Create với các bản ghi mới
	return r.db.Create(&newSchedulesToCreate).Error
}

func (r *shiftScheduleRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.ShiftSchedule{}).Where("id = ?", id).Updates(updates).Error
}

func (r *shiftScheduleRepo) Delete(id uint) error {
	return r.db.Delete(&model.ShiftSchedule{}, id).Error
}
