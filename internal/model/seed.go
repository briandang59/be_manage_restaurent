package model

import (
	"gorm.io/gorm"
	"log"
)

func SeedRolesAndPermissions(db *gorm.DB) {
	// Seed permissions
	permissions := []Permission{
		// Quản lý tài khoản
		{Name: "view_accounts"}, {Name: "create_account"}, {Name: "edit_account"}, {Name: "delete_account"},
		// Quản lý nhân viên
		{Name: "view_employees"}, {Name: "create_employee"}, {Name: "edit_employee"}, {Name: "delete_employee"},
		// Quản lý ca làm việc
		{Name: "view_shifts"}, {Name: "create_shift"}, {Name: "edit_shift"}, {Name: "delete_shift"},
		// Quản lý lịch làm việc
		{Name: "view_shift_schedules"}, {Name: "create_shift_schedule"}, {Name: "edit_shift_schedule"}, {Name: "delete_shift_schedule"},
		// Quản lý thực đơn
		{Name: "view_menu_items"}, {Name: "create_menu_item"}, {Name: "edit_menu_item"}, {Name: "delete_menu_item"},
		// Quản lý nguyên liệu
		{Name: "view_ingredients"}, {Name: "create_ingredient"}, {Name: "edit_ingredient"}, {Name: "delete_ingredient"},
		// Quản lý kho
		{Name: "view_inventory"}, {Name: "edit_inventory"},
		// Quản lý đơn hàng
		{Name: "view_orders"}, {Name: "create_order"}, {Name: "edit_order"}, {Name: "delete_order"},
		// Quản lý bàn
		{Name: "view_tables"}, {Name: "create_table"}, {Name: "edit_table"}, {Name: "delete_table"},
		// Quản lý khách hàng
		{Name: "view_customers"}, {Name: "create_customer"}, {Name: "edit_customer"}, {Name: "delete_customer"},
		// Quản lý quyền/vai trò
		{Name: "view_roles"}, {Name: "create_role"}, {Name: "edit_role"}, {Name: "delete_role"}, {Name: "view_permissions"},
		// Quản lý phiếu nhập/xuất kho
		{Name: "view_tickets"}, {Name: "create_ticket"}, {Name: "edit_ticket"}, {Name: "delete_ticket"},
		// Chấm công
		{Name: "view_attendance"}, {Name: "create_attendance"}, {Name: "edit_attendance"}, {Name: "delete_attendance"},
		// Quản lý file
		{Name: "view_files"}, {Name: "upload_file"}, {Name: "delete_file"},
	}
	for _, p := range permissions {
		if err := db.FirstOrCreate(&p, Permission{Name: p.Name}).Error; err != nil {
			log.Println("Seed permission error:", err)
		}
	}

	// Seed roles
	roles := []Role{
		{RoleName: "Admin"},
		{RoleName: "Manager"},
		{RoleName: "Cashier"},
		{RoleName: "Waiter"},
		{RoleName: "Chef"},
		{RoleName: "KitchenStaff"},
		{RoleName: "InventoryStaff"},
		{RoleName: "Customer"},
	}
	for _, r := range roles {
		if err := db.FirstOrCreate(&r, Role{RoleName: r.RoleName}).Error; err != nil {
			log.Println("Seed role error:", err)
		}
	}
} 