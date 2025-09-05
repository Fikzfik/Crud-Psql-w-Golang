package service

import (
	"crud-alumni/model"
	"crud-alumni/repository"
)

// Ambil semua pekerjaan alumni
func GetAllPekerjaan() ([]model.PekerjaanAlumni, error) {
	return repository.GetAllPekerjaan()
}

// Ambil pekerjaan by ID
func GetPekerjaanByID(id int) (model.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByID(id)
}

// Ambil pekerjaan berdasarkan alumni
func GetPekerjaanByAlumni(alumniID int) ([]model.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByAlumni(alumniID)
}

// Tambah pekerjaan
func CreatePekerjaan(p model.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.InsertPekerjaan(p)
}

// Update pekerjaan
func UpdatePekerjaan(id int, p model.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.UpdatePekerjaan(id, p)
}

// Hapus pekerjaan
func DeletePekerjaan(id int) error {
	return repository.DeletePekerjaan(id)
}


