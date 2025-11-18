package repository

import "crud-alumni/app/models"
type AlumniRepository interface {
	GetAlumniByName(name string) ([]models.Alumni, error)
	InsertAlumni(a models.Alumni) error
	GetAlumniByID(id string) (models.Alumni, error)
	GetAllAlumni() ([]models.Alumni, error)
	UpdateAlumni(id string, a models.Alumni) (models.Alumni, error)
	DeleteAlumni(id string) error
}
