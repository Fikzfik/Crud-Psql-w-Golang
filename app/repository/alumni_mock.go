package repository

import (
	"crud-alumni/app/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockAlumniRepo struct {
	Data map[string]models.Alumni
}

func NewMockAlumniRepo() *MockAlumniRepo {
	return &MockAlumniRepo{
		Data: make(map[string]models.Alumni),
	}
}

func (m *MockAlumniRepo) InsertAlumni(a models.Alumni) error {
	if a.Nama == "" || a.NIM == "" {
		return errors.New("invalid data")
	}

	if a.ID == primitive.NilObjectID {
		a.ID = primitive.NewObjectID()
	}

	m.Data[a.ID.Hex()] = a
	return nil
}

func (m *MockAlumniRepo) GetAlumniByID(id string) (models.Alumni, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Alumni{}, errors.New("invalid ID")
	}

	result, ok := m.Data[objID.Hex()]
	if !ok {
		return models.Alumni{}, errors.New("not found")
	}
	return result, nil
}
func (m *MockAlumniRepo) GetAlumniByName(name string) ([]models.Alumni, error) {
	if name == "" {
		return []models.Alumni{}, errors.New("invalid name")
	}

	var results []models.Alumni
	for _, v := range m.Data {
		if v.Nama == name {
			results = append(results, v)
		}
	}

	if len(results) == 0 {
		return []models.Alumni{}, errors.New("not found")
	}

	return results, nil
}
func (m *MockAlumniRepo) GetAllAlumni() ([]models.Alumni, error) {
	list := []models.Alumni{}
	for _, v := range m.Data {
		list = append(list, v)
	}
	return list, nil
}

func (m *MockAlumniRepo) UpdateAlumni(id string, a models.Alumni) (models.Alumni, error) {
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Alumni{}, errors.New("invalid ID")
	}

	_, exists := m.Data[id]
	if !exists {
		return models.Alumni{}, errors.New("not found")
	}

	if a.Nama == "" {
		return models.Alumni{}, errors.New("invalid nama")
	}
	if a.NIM == "" {
		return models.Alumni{}, errors.New("invalid NIM")
	}

	a.ID, _ = primitive.ObjectIDFromHex(id)
	m.Data[id] = a

	return a, nil
}

func (m *MockAlumniRepo) DeleteAlumni(id string) error {
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID")
	}

	_, exists := m.Data[id]
	if !exists {
		return errors.New("not found")
	}

	delete(m.Data, id)
	return nil
}
