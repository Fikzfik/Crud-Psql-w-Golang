package service

import (
	"crud-alumni/model"
	"crud-alumni/repository"
)

// Ambil semua data alumni
func GetAllAlumni() ([]model.Alumni, error) {
	return repository.GetAllAlumni()
}

// Ambil 1 alumni berdasarkan ID
func GetAlumniByID(id int) (model.Alumni, error) {
	return repository.GetAlumniByID(id)
}

// Tambah alumni baru
func CreateAlumni(a model.Alumni) error {
	// validasi sederhana
	if a.NIM == "" || a.Nama == "" || a.Email == "" {
		return ErrInvalidData
	}
	return repository.InsertAlumni(a)
}

// Update data alumni
func UpdateAlumni(id int, a model.Alumni) error {
	if a.NIM == "" || a.Nama == "" || a.Email == "" {
		return ErrInvalidData
	}
	return repository.UpdateAlumni(id, a)
}

// Hapus alumni
func DeleteAlumni(id int) error {
	return repository.DeleteAlumni(id)
}
