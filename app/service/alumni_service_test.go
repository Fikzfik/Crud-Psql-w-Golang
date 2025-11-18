package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateAlumni_Success(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	a := models.Alumni{
		Nama: "Fikri",
		NIM:  "A001",
	}

	err := mock.InsertAlumni(a)
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}

	if len(mock.Data) != 1 {
		t.Errorf("Expected 1 data, got %d", len(mock.Data))
	}
}

func TestCreateAlumni_EmptyNama(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	a := models.Alumni{
		Nama: "",
		NIM:  "A001",
	}

	err := mock.InsertAlumni(a)
	if err == nil {
		t.Errorf("Expected error for empty nama, got nil")
	}
}

func TestCreateAlumni_EmptyNIM(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	a := models.Alumni{
		Nama: "Test",
		NIM:  "",
	}

	err := mock.InsertAlumni(a)
	if err == nil {
		t.Errorf("Expected error for empty NIM, got nil")
	}
}

//
// GET BY ID
//

func TestGetAlumniByID_Success(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	id := primitive.NewObjectID()
	a := models.Alumni{
		ID:   id,
		Nama: "Fikri",
		NIM:  "A001",
	}

	mock.InsertAlumni(a)

	result, err := mock.GetAlumniByID(id.Hex())
	if err != nil {
		t.Errorf("Expected success, got error %v", err)
	}

	if result.Nama != "Fikri" {
		t.Errorf("Expected 'Fikri', got %s", result.Nama)
	}
}

func TestGetAlumniByID_NotFound(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	_, err := mock.GetAlumniByID(primitive.NewObjectID().Hex())
	if err == nil {
		t.Errorf("Expected error when alumni not found")
	}
}

func TestGetAlumniByID_InvalidID(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	_, err := mock.GetAlumniByID("invalid_hex")
	if err == nil {
		t.Errorf("Expected error for invalid hex ID")
	}
}

//
// GET ALL
//

func TestGetAllAlumni(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	mock.InsertAlumni(models.Alumni{ID: primitive.NewObjectID(), Nama: "A"})
	mock.InsertAlumni(models.Alumni{ID: primitive.NewObjectID(), Nama: "B"})

	all, err := mock.GetAllAlumni()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(all) != 2 {
		t.Errorf("Expected 2 alumni, got %d", len(all))
	}
}

//
// UPDATE
//

func TestUpdateAlumni_Success(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	id := primitive.NewObjectID()
	mock.InsertAlumni(models.Alumni{
		ID:   id,
		Nama: "Old Name",
		NIM:  "A001",
	})

	newData := models.Alumni{
		Nama: "New Name",
		NIM:  "A001",
	}

	updated, err := mock.UpdateAlumni(id.Hex(), newData)
	if err != nil {
		t.Errorf("Update error: %v", err)
	}

	if updated.Nama != "New Name" {
		t.Errorf("Expected 'New Name', got '%s'", updated.Nama)
	}
}

func TestUpdateAlumni_NotFound(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	_, err := mock.UpdateAlumni(primitive.NewObjectID().Hex(), models.Alumni{
		Nama: "x",
		NIM:  "y",
	})

	if err == nil {
		t.Errorf("Expected error for not found")
	}
}

func TestUpdateAlumni_EmptyNama(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	id := primitive.NewObjectID()
	mock.InsertAlumni(models.Alumni{
		ID:   id,
		Nama: "Old",
		NIM:  "A001",
	})

	_, err := mock.UpdateAlumni(id.Hex(), models.Alumni{
		Nama: "",
		NIM:  "A001",
	})

	if err == nil {
		t.Errorf("Expected error for empty name")
	}
}

//
// DELETE
//

func TestDeleteAlumni_Success(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	id := primitive.NewObjectID()
	mock.InsertAlumni(models.Alumni{ID: id, Nama: "DeleteMe", NIM: "A001"})

	err := mock.DeleteAlumni(id.Hex())
	if err != nil {
		t.Errorf("Delete error: %v", err)
	}

	if len(mock.Data) != 0 {
		t.Errorf("Expected data empty after delete, got %d", len(mock.Data))
	}
}

func TestDeleteAlumni_NotFound(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	err := mock.DeleteAlumni(primitive.NewObjectID().Hex())
	if err == nil {
		t.Errorf("Expected not found error")
	}
}

func TestDeleteAlumni_InvalidID(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	err := mock.DeleteAlumni("invalid_hex")
	if err == nil {
		t.Errorf("Expected invalid ID error")
	}
}

//
// STRESS TEST
//

func TestInsertManyAlumni(t *testing.T) {
	mock := repository.NewMockAlumniRepo()

	for i := 0; i < 50; i++ {
		a := models.Alumni{
			Nama: "User",
			NIM:  primitive.NewObjectID().Hex(),
		}
		_ = mock.InsertAlumni(a)
	}

	if len(mock.Data) != 50 {
		t.Errorf("Expected 50 data, got %d", len(mock.Data))
	}
}
