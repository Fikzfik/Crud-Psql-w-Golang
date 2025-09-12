package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
)

// Ambil semua data alumni
func GetAllAlumniByFak(fak string, a models.Alumni) ([]models.Alumni, error) {
	if fak == "" {
		return nil,ErrInvalidData
	}else{
		return repository.GetAllAlumniByFak(fak, a)
	}
}

func GetAllAlumni()([]models.Alumni,error){
	return repository.GetAllAlumni()
}
// Ambil 1 alumni berdasarkan ID
func GetAlumniByID(id int) (models.Alumni, error) {

	return repository.GetAlumniByID(id)
}

// Tambah alumni baru
func CreateAlumni(a models.Alumni) error {
	// validasi sederhana
	if a.NIM == "" || a.Nama == "" || a.Email == "" {
		return ErrInvalidData
	}
	return repository.InsertAlumni(a)
}

// Update data alumni
func UpdateAlumni(id int, a models.Alumni) error {
	if a.NIM == "" || a.Nama == "" || a.Email == "" {
		return ErrInvalidData
	}
	return repository.UpdateAlumni(id, a)
}

// Hapus alumni
func DeleteAlumni(id int) error {
	return repository.DeleteAlumni(id)
}
