package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import "time"

// Struct Alumni sesuai collection alumni
type Alumni struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	NIM        string             `bson:"nim" json:"nim"`
	Nama       string             `bson:"nama" json:"nama"`
	Jurusan    string             `bson:"jurusan" json:"jurusan"`
	Fakultas   string             `bson:"fakultas" json:"fakultas"`
	Angkatan   int                `bson:"angkatan" json:"angkatan"`
	TahunLulus int                `bson:"tahun_lulus" json:"tahun_lulus"`
	NoTelepon  string             `bson:"no_telepon" json:"no_telepon"`
	Alamat     string             `bson:"alamat" json:"alamat"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
