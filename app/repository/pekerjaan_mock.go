package repository

import (
	"crud-alumni/app/models"
	"errors"
)

type MockPekerjaanRepo struct {
	Data map[string]models.PekerjaanAlumni
}

func NewMockPekerjaanRepo() *MockPekerjaanRepo {
	return &MockPekerjaanRepo{
		Data: make(map[string]models.PekerjaanAlumni),
	}
}

func (m *MockPekerjaanRepo) InsertPekerjaan(p models.PekerjaanAlumni) error {
	m.Data[p.ID.Hex()] = p
	return nil
}

func (m *MockPekerjaanRepo) GetPekerjaanByID(id string) (models.PekerjaanAlumni, error) {
	if val, ok := m.Data[id]; ok {
		return val, nil
	}
	return models.PekerjaanAlumni{}, errors.New("not found")
}

func (m *MockPekerjaanRepo) GetPekerjaanByAlumni(alumniID string) ([]models.PekerjaanAlumni, error) {
	var result []models.PekerjaanAlumni
	for _, p := range m.Data {
		if p.AlumniID.Hex() == alumniID {
			result = append(result, p)
		}
	}
	return result, nil
}

func (m *MockPekerjaanRepo) UpdatePekerjaan(id string, p models.PekerjaanAlumni) (models.PekerjaanAlumni, error) {
	if _, ok := m.Data[id]; !ok {
		return models.PekerjaanAlumni{}, errors.New("not found")
	}
	m.Data[id] = p
	return p, nil
}

func (m *MockPekerjaanRepo) DeletePekerjaan(id string) error {
	if _, ok := m.Data[id]; !ok {
		return errors.New("not found")
	}
	delete(m.Data, id)
	return nil
}
