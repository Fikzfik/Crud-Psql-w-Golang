package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
)

// Ambil semua pekerjaan alumni
func GetAllPekerjaan() ([]models.PekerjaanAlumni, error) {
	return repository.GetAllPekerjaan()
}

// Ambil pekerjaan by ID
func GetPekerjaanByID(id int) (models.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByID(id)
}

// Ambil pekerjaan berdasarkan alumni
func GetPekerjaanByAlumni(alumniID int) ([]models.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByAlumni(alumniID)
}

// Tambah pekerjaan
func CreatePekerjaan(p models.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.InsertPekerjaan(p)
}

// Update pekerjaan
func UpdatePekerjaan(id int, p models.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.UpdatePekerjaan(id, p)
}

// Hapus pekerjaan
func DeletePekerjaan(id int) error {
	return repository.DeletePekerjaan(id)
}

func IsDeleted(role string,id int) error {
	return repository.IsDeleted(role,id)
}
