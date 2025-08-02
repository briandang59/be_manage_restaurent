package handler

import (
	"context"
	"fmt"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/repository"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	repo *repository.FileRepo
}

func NewFileHandler(repo *repository.FileRepo) *FileHandler {
	return &FileHandler{repo: repo}
}

// UploadFile godoc
// @Summary Upload file
// @Description Upload file (hình ảnh, tài liệu...)
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File upload"
// @Success 200 {object} model.File
// @Failure 400 {object} response.ErrorResponse
// @Router /files/upload [post]
func (h *FileHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Không có file upload")
		return
	}

	f, err := file.Open()
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Không mở được file")
		return
	}
	defer f.Close()

	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Cloudinary config error")
		return
	}

	// Tạo PublicID duy nhất với timestamp
	timestamp := time.Now().Unix()
	uniquePublicID := fmt.Sprintf("uploads/%d_%s", timestamp, file.Filename)
	
	fmt.Printf("Uploading file: %s with PublicID: %s\n", file.Filename, uniquePublicID)
	
	uploadResult, err := cld.Upload.Upload(context.Background(), f, uploader.UploadParams{
		Folder: "uploads",
		PublicID: uniquePublicID,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, fmt.Sprintf("Upload lỗi: %v", err))
		return
	}
	
	fmt.Printf("Upload result: PublicID=%s, URL=%s\n", uploadResult.PublicID, uploadResult.SecureURL)

	fileModel := model.File{
		FileName:     file.Filename,
		Url:          uploadResult.SecureURL,
		MimeType:     file.Header.Get("Content-Type"),
		Size:         file.Size,
		PublicID:     uploadResult.PublicID,
		ResourceType: uploadResult.ResourceType,
	}
	if err := h.repo.Create(&fileModel); err != nil {
		response.Error(c, http.StatusInternalServerError, "Lưu file vào DB lỗi")
		return
	}
	response.Success(c, fileModel, nil)
}

// GetFile godoc
// @Summary Lấy thông tin file
// @Description Lấy thông tin file theo ID
// @Tags file
// @Produce json
// @Param id path int true "ID file"
// @Success 200 {object} model.File
// @Failure 404 {object} response.ErrorResponse
// @Router /files/{id} [get]
func (h *FileHandler) GetFile(c *gin.Context) {
	// ... code lấy file ...
	response.Success(c, model.File{ID: 1, FileName: "example.jpg", Url: "https://example.com/example.jpg"}, nil)
}

// DeleteFile godoc
// @Summary Xóa file
// @Description Xóa một file theo ID
// @Tags file
// @Produce json
// @Param id path int true "ID file"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /files/{id} [delete]
func (h *FileHandler) DeleteFile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID không hợp lệ")
		return
	}

	// Lấy thông tin file từ database
	file, err := h.repo.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Không tìm thấy file")
		return
	}

	fmt.Printf("Đang xóa file: ID=%d, PublicID=%s, FileName=%s\n", file.ID, file.PublicID, file.FileName)

	// Kiểm tra PublicID có tồn tại không
	if file.PublicID == "" {
		fmt.Printf("PublicID rỗng, bỏ qua xóa trên Cloudinary\n")
	} else {
		// Xóa file trên Cloudinary
		cld, err := cloudinary.NewFromParams(
			os.Getenv("CLOUDINARY_CLOUD_NAME"),
			os.Getenv("CLOUDINARY_API_KEY"),
			os.Getenv("CLOUDINARY_API_SECRET"),
		)
		if err != nil {
			fmt.Printf("Lỗi khởi tạo Cloudinary: %v\n", err)
			response.Error(c, http.StatusInternalServerError, "Cloudinary config error")
			return
		}

		fmt.Printf("Cloudinary config: CloudName=%s, APIKey=%s\n", 
			os.Getenv("CLOUDINARY_CLOUD_NAME"), 
			os.Getenv("CLOUDINARY_API_KEY"))

		// Xóa file trên Cloudinary
		result, err := cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
			PublicID: file.PublicID,
		})
		if err != nil {
			fmt.Printf("Lỗi xóa file trên Cloudinary: %v\n", err)
			fmt.Printf("PublicID được sử dụng: %s\n", file.PublicID)
		} else {
			fmt.Printf("Xóa file trên Cloudinary thành công: %+v\n", result)
		}
	}

	// Xóa record trong database
	if err := h.repo.Delete(uint(id)); err != nil {
		fmt.Printf("Lỗi xóa record trong database: %v\n", err)
		response.Error(c, http.StatusInternalServerError, "Lỗi xóa file trong database")
		return
	}

	fmt.Printf("Xóa record trong database thành công\n")
	response.Success(c, "File deleted", nil)
}

// ListFiles godoc
// @Summary Lấy danh sách file
// @Description Lấy danh sách file đã upload
// @Tags file
// @Produce json
// @Success 200 {object} []model.File
// @Router /files [get]
func (h *FileHandler) ListFiles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	files, total, err := h.repo.List(offset, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Không lấy được danh sách file")
		return
	}
	response.Success(c, files, &response.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
	})
} 