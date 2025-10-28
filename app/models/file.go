package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type File struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"` // 🔥 tambahkan ini
	FileName     string             `bson:"file_name" json:"file_name"`
	OriginalName string             `bson:"original_name" json:"original_name"`
	FilePath     string             `bson:"file_path" json:"file_path"`
	FileSize     int64              `bson:"file_size" json:"file_size"`
	FileType     string             `bson:"file_type" json:"file_type"`
	UploadedAt   time.Time          `bson:"uploaded_at" json:"uploaded_at"`
}
type FileResponse struct {
	ID           string    `json:"id"`
	FileName     string    `json:"file_name"`
	OriginalName string    `json:"original_name"`
	FilePath     string    `json:"file_path"`
	FileSize     int64     `json:"file_size"`
	FileType     string    `json:"file_type"`
	UploadedAt   time.Time `json:"uploaded_at"`
}
