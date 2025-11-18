package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ===== Upload Foto =====
// UploadFoto godoc
// @Summary Upload foto pengguna
// @Description Upload foto profil untuk user. Hanya admin bisa upload foto milik user lain.
// @Tags Upload
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param foto formData file true "File foto (JPEG, JPG, PNG, max 1MB)"
// @Param user_id formData string false "ID user (opsional, hanya admin boleh isi)"
// @Success 201 {object} models.FileResponse
// @Failure 400 {object} map[string]interface{} "Request tidak valid"
// @Failure 403 {object} map[string]interface{} "Tidak diizinkan upload user lain"
// @Failure 500 {object} map[string]interface{} "Gagal menyimpan file atau metadata"
// @Router /foto/upload [post]
func UploadFoto(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	fmt.Println("DEBUG FORM:", form, err)
	fileHeader, err := c.FormFile("foto")
	fmt.Println("DEBUG FILE:", fileHeader, err)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Foto tidak ditemukan"})
	}

	// ===== Role dan User ID dari JWT Middleware =====
	role := c.Locals("role").(string)
	userID := c.Locals("user_id").(string)
	targetID := c.FormValue("user_id")

	// Jika tidak mengirim user_id, otomatis pakai ID miliknya sendiri
	if targetID == "" {
		targetID = userID
	}

	// Jika bukan admin dan mencoba upload untuk user lain
	if role != "admin" && targetID != userID {
		return c.Status(403).JSON(fiber.Map{"message": "User tidak diizinkan upload foto untuk orang lain"})
	}

	// Konversi targetID dari string ke ObjectID
	objUserID, err := primitive.ObjectIDFromHex(targetID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "User ID tidak valid"})
	}

	// ===== Validasi ukuran dan tipe =====
	if fileHeader.Size > 1*1024*1024 {
		return c.Status(400).JSON(fiber.Map{"message": "Ukuran foto maksimal 1MB"})
	}

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		return c.Status(400).JSON(fiber.Map{"message": "Format foto tidak diperbolehkan"})
	}

	// ===== Simpan file =====
	uploadPath := "./uploads/foto"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal membuat folder upload"})
	}

	ext := filepath.Ext(fileHeader.Filename)
	newName := uuid.New().String() + ext
	savePath := filepath.Join(uploadPath, newName)

	if err := saveFile(fileHeader, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal menyimpan foto"})
	}

	// ===== Simpan metadata ke DB =====
	data := models.File{
		UserID:       objUserID,
		FileName:     newName,
		OriginalName: fileHeader.Filename,
		FilePath:     savePath,
		FileSize:     fileHeader.Size,
		FileType:     contentType,
		UploadedAt:   time.Now(),
	}

	if err := repository.InsertFoto(data); err != nil {
		os.Remove(savePath)
		return c.Status(500).JSON(fiber.Map{"message": "Gagal simpan metadata ke DB"})
	}

	// Buat respons
	response := models.FileResponse{
		ID:           data.ID.Hex(),
		FileName:     data.FileName,
		OriginalName: data.OriginalName,
		FilePath:     data.FilePath,
		FileSize:     data.FileSize,
		FileType:     data.FileType,
		UploadedAt:   data.UploadedAt,
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Foto berhasil diupload",
		"data":    response,
	})
}

// ===== Upload Sertifikat =====
// UploadSertifikat godoc
// @Summary Upload sertifikat pengguna
// @Description Upload file sertifikat dalam format PDF (hanya admin bisa upload milik user lain)
// @Tags Upload
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param sertifikat formData file true "File sertifikat (PDF, max 2MB)"
// @Param user_id formData string false "ID user (opsional, hanya admin boleh isi)"
// @Success 201 {object} models.FileResponse
// @Failure 400 {object} map[string]interface{} "Request tidak valid"
// @Failure 403 {object} map[string]interface{} "Tidak diizinkan upload user lain"
// @Failure 500 {object} map[string]interface{} "Gagal menyimpan file atau metadata"
// @Router /sertifikat/upload [post]
func UploadSertifikat(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("sertifikat")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Sertifikat tidak ditemukan"})
	}

	if fileHeader.Size > 2*1024*1024 {
		return c.Status(400).JSON(fiber.Map{"message": "Ukuran sertifikat maksimal 2MB"})
	}

	if fileHeader.Header.Get("Content-Type") != "application/pdf" {
		return c.Status(400).JSON(fiber.Map{"message": "Format sertifikat harus PDF"})
	}

	// ===== Ambil role dan user dari JWT =====
	role := c.Locals("role").(string)
	userID := c.Locals("user_id").(string)
	targetID := c.FormValue("user_id")

	// Jika tidak mengirim user_id → dianggap upload dirinya sendiri
	if targetID == "" {
		targetID = userID
	}

	// Jika bukan admin dan upload milik orang lain
	if role != "admin" && targetID != userID {
		return c.Status(403).JSON(fiber.Map{"message": "User hanya bisa upload sertifikat miliknya sendiri"})
	}

	// Konversi targetID ke ObjectID
	objUserID, err := primitive.ObjectIDFromHex(targetID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "User ID tidak valid"})
	}

	uploadPath := "./uploads/sertifikat"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal membuat folder upload"})
	}

	ext := filepath.Ext(fileHeader.Filename)
	newName := uuid.New().String() + ext
	savePath := filepath.Join(uploadPath, newName)

	if err := saveFile(fileHeader, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal menyimpan sertifikat"})
	}

	data := models.File{
		UserID:       objUserID,
		FileName:     newName,
		OriginalName: fileHeader.Filename,
		FilePath:     savePath,
		FileSize:     fileHeader.Size,
		FileType:     "application/pdf",
		UploadedAt:   time.Now(),
	}

	if err := repository.InsertSertifikat(data); err != nil {
		os.Remove(savePath)
		return c.Status(500).JSON(fiber.Map{"message": "Gagal simpan metadata ke DB"})
	}

	response := models.FileResponse{
		ID:           data.ID.Hex(),
		FileName:     data.FileName,
		OriginalName: data.OriginalName,
		FilePath:     data.FilePath,
		FileSize:     data.FileSize,
		FileType:     data.FileType,
		UploadedAt:   data.UploadedAt,
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Sertifikat berhasil diupload",
		"data":    response,
	})
}

// ===== Helper =====
func saveFile(fileHeader *multipart.FileHeader, path string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(file)
	return err
}
