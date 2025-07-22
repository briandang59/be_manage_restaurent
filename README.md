# Restaurant Management REST API

## 1. Yêu cầu hệ thống
- Go 1.18+
- MySQL/PostgreSQL (hoặc DB tương thích GORM)
- (Tùy chọn) Docker, Docker Compose

## 2. Cài đặt & cấu hình

### Clone project
```bash
git clone <repo-url>
cd be_manage_restaurent
```

### Cài đặt package (nếu chạy local)
```bash
go mod tidy
```

### Cấu hình biến môi trường
Tạo file `.env` với nội dung mẫu:
```
# DB cho local hoặc Docker Compose
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=restaurant_db
DB_DSN="host=localhost user=postgres password=postgres dbname=restaurant_db port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

# Backend
BE_PORT=8080

# Cloudinary
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
```

## 3. Chạy project

### Chạy local
```bash
go run cmd/server/main.go
```
- Server mặc định chạy tại: [http://localhost:8080](http://localhost:8080)

### Chạy bằng Docker Compose
```bash
docker-compose up --build
```
- Backend mặc định chạy tại: [http://localhost:5123](http://localhost:5123) (xem lại port mapping trong docker-compose.yml)
- Database Postgres: [localhost:5432]
- Đảm bảo file `.env` đúng và đủ biến môi trường cho cả DB, Cloudinary, backend.

## 4. Swagger API Docs
- Truy cập [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) (hoặc port bạn cấu hình) để xem tài liệu API tự động.
- Đầy đủ mô tả các trường, dữ liệu mẫu, response mẫu cho tất cả endpoint.

## 5. Database migration
- Khi chạy lần đầu, project sẽ tự động migrate các bảng cần thiết (AutoMigrate).
- Nếu thay đổi model, hãy chạy lại project để cập nhật DB.

## 6. Upload file với Cloudinary
- Endpoint: `POST /api/files/upload` (multipart/form-data, param `file`)
- File sẽ được upload lên Cloudinary, thông tin file (url, public_id, ...) sẽ được lưu vào DB.
- Để xóa file, dùng endpoint: `DELETE /api/files/{id}`

## 7. Các endpoint chính
- CRUD cho: Account, Role, Permission, Employee, Customer, MenuItem, Table, Order, OrderItem, Ingredient, Attendance, Shift, ShiftSchedule, File...
- Đăng nhập: `POST /api/accounts/login` (trả về JWT token)
- Xem chi tiết endpoint và dữ liệu mẫu tại Swagger.

## 8. Một số lưu ý
- Đảm bảo DB đã khởi động trước khi chạy project (nếu dùng Docker Compose, DB sẽ tự chạy cùng app).
- Để upload file lên Cloudinary, cần cấu hình đúng các biến môi trường Cloudinary.
- Có thể mở rộng thêm các tính năng phân quyền, caching, logging, ...
- Nếu đổi port backend, hãy sửa lại port trong Dockerfile, docker-compose.yml và biến môi trường cho đồng bộ.

## 9. Đóng góp & phát triển
- Fork, tạo branch, pull request như các dự án open source thông thường.
- Đóng góp thêm test, validate, CI/CD, ... rất hoan nghênh!

---
**Mọi thắc mắc hoặc lỗi, vui lòng liên hệ hoặc tạo issue trên repo.** 