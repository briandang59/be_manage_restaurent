package dto

import "time"

type CreateAttendanceDTO struct {
	ShiftScheduleId uint      `json:"shift_schedule_id" binding:"required"`
	ActualStartTime time.Time `json:"actual_start_time" binding:"required"`
	ActualEndTime   time.Time `json:"actual_end_time" binding:"required"`
	// Hours sẽ được tự động tính toán từ ActualStartTime và ActualEndTime
}

type UpdateAttendanceDTO struct {
	ActualStartTime *time.Time `json:"actual_start_time,omitempty"`
	ActualEndTime   *time.Time `json:"actual_end_time,omitempty"`
	// Hours sẽ được tự động tính toán lại nếu cập nhật thời gian
}
