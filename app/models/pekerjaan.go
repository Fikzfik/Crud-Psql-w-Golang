package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct PekerjaanAlumni sesuai collection pekerjaan_alumni
type PekerjaanAlumni struct {
	ID                 primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	AlumniID           primitive.ObjectID  `bson:"alumni_id,omitempty" json:"alumni_id,omitempty"`
	NamaPerusahaan     string              `bson:"nama_perusahaan" json:"nama_perusahaan"`
	PosisiJabatan      string              `bson:"posisi_jabatan" json:"posisi_jabatan"`
	BidangIndustri     string              `bson:"bidang_industri" json:"bidang_industri"`
	LokasiKerja        string              `bson:"lokasi_kerja" json:"lokasi_kerja"`
	GajiRange          string              `bson:"gaji_range" json:"gaji_range"`
	TanggalMulaiKerja  time.Time           `bson:"tanggal_mulai_kerja" json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *time.Time          `bson:"tanggal_selesai_kerja,omitempty" json:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan    string              `bson:"status_pekerjaan" json:"status_pekerjaan"`
	DeskripsiPekerjaan string              `bson:"deskripsi_pekerjaan" json:"deskripsi_pekerjaan"`
	IsDeleted          bool                `bson:"is_deleted" json:"is_deleted"`
	CreatedAt          time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt          time.Time           `bson:"updated_at" json:"updated_at"`
}
