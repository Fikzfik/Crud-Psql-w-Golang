package repository

import "crud-alumni/app/models"

type PekerjaanRepository interface {
	GetPekerjaanByID(id string) (models.PekerjaanAlumni, error)
	GetPekerjaanByAlumni(alumniID string) ([]models.PekerjaanAlumni, error)
	InsertPekerjaan(p models.PekerjaanAlumni) error
	UpdatePekerjaan(id string, p models.PekerjaanAlumni) (models.PekerjaanAlumni, error)
	DeletePekerjaan(id string) error
}
