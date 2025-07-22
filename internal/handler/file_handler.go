package handler

import (
	"context"
	"fmt"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type FileHandler struct{}

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

	uploadResult, err := cld.Upload.Upload(context.Background(), f, uploader.UploadParams{
		Folder: "uploads",
		PublicID: file.Filename,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, fmt.Sprintf("Upload lỗi: %v", err))
		return
	}

	fileModel := model.File{
		FileName:     file.Filename,
		Url:          uploadResult.SecureURL,
		MimeType:     file.Header.Get("Content-Type"),
		Size:         file.Size,
		PublicID:     uploadResult.PublicID,
		ResourceType: uploadResult.ResourceType,
		Folder:       uploadResult.Folder,
	}
	// Lưu vào DB (giả sử có h.repo kiểu repository.FileRepo)
	if h.repo != nil {
		err = h.repo.Create(&fileModel)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Lưu file vào DB lỗi")
			return
		}
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
	// ... code xóa file ...
	response.Success(c, "File deleted", nil)
} 