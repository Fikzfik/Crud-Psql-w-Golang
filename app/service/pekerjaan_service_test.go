package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreatePekerjaan(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	p := models.PekerjaanAlumni{
		ID:             primitive.NewObjectID(),
		NamaPerusahaan: "PT ABC",
		PosisiJabatan:  "Backend Engineer",
	}

	err := mock.InsertPekerjaan(p)
	if err != nil {
		t.Errorf("Expected success, got error %v", err)
	}

	if len(mock.Data) != 1 {
		t.Errorf("Expected 1 data, got %d", len(mock.Data))
	}
}

func TestCreatePekerjaan_EmptyCompany(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	p := models.PekerjaanAlumni{
		ID:             primitive.NewObjectID(),
		NamaPerusahaan: "",
		PosisiJabatan:  "Engineer",
	}

	err := mock.InsertPekerjaan(p)
	if err == nil {
		t.Errorf("Expected error for empty NamaPerusahaan, got nil")
	}
}

func TestCreatePekerjaan_EmptyJabatan(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	p := models.PekerjaanAlumni{
		ID:             primitive.NewObjectID(),
		NamaPerusahaan: "PT ABC",
		PosisiJabatan:  "",
	}

	err := mock.InsertPekerjaan(p)
	if err == nil {
		t.Errorf("Expected error for empty PosisiJabatan, got nil")
	}
}
func TestCreatePekerjaan_DuplicateID(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	id := primitive.NewObjectID()
	p1 := models.PekerjaanAlumni{ID: id}
	p2 := models.PekerjaanAlumni{ID: id}

	_ = mock.InsertPekerjaan(p1)
	err := mock.InsertPekerjaan(p2)

	if err == nil {
		t.Errorf("Expected error for duplicate ID, got nil")
	}
}

func TestGetPekerjaanByID_InvalidID(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	_, err := mock.GetPekerjaanByID("INVALID_HEX")
	if err == nil {
		t.Errorf("Expected error for invalid hex, got nil")
	}
}

func TestGetPekerjaanByID(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	id := primitive.NewObjectID()
	p := models.PekerjaanAlumni{
		ID:             id,
		NamaPerusahaan: "PT XYZ",
	}

	mock.InsertPekerjaan(p)

	result, err := mock.GetPekerjaanByID(id.Hex())
	if err != nil {
		t.Errorf("Expected success, got error %v", err)
	}

	if result.NamaPerusahaan != "PT XYZ" {
		t.Errorf("Expected PT XYZ, got %s", result.NamaPerusahaan)
	}
}


func TestGetPekerjaanByAlumni(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	alumniID := primitive.NewObjectID()
	p := models.PekerjaanAlumni{
		ID:       primitive.NewObjectID(),
		AlumniID: alumniID,
	}

	mock.InsertPekerjaan(p)

	data, err := mock.GetPekerjaanByAlumni(alumniID.Hex())
	if err != nil {
		t.Errorf("Expected success, got %v", err)
	}

	if len(data) != 1 {
		t.Errorf("Expected 1 result, got %d", len(data))
	}
}
func TestGetPekerjaanByAlumni_Empty(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	data, _ := mock.GetPekerjaanByAlumni(primitive.NewObjectID().Hex())
	if len(data) != 0 {
		t.Errorf("Expected 0 pekerjaan, got %d", len(data))
	}
}

func TestUpdatePekerjaan(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	id := primitive.NewObjectID()
	p := models.PekerjaanAlumni{
		ID:             id,
		NamaPerusahaan: "PT Lama",
	}

	mock.InsertPekerjaan(p)

	newData := models.PekerjaanAlumni{
		ID:             id,
		NamaPerusahaan: "PT Baru",
	}

	updated, err := mock.UpdatePekerjaan(id.Hex(), newData)
	if err != nil {
		t.Errorf("Update error: %v", err)
	}

	if updated.NamaPerusahaan != "PT Baru" {
		t.Errorf("Expected PT Baru, got %s", updated.NamaPerusahaan)
	}
}
func TestGetPekerjaanByAlumni_InvalidID(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	_, err := mock.GetPekerjaanByAlumni("bad_id")
	if err == nil {
		t.Errorf("Expected error for invalid alumni ID, got nil")
	}
}

func TestDeletePekerjaan(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	id := primitive.NewObjectID()
	p := models.PekerjaanAlumni{
		ID: id,
	}

	mock.InsertPekerjaan(p)

	err := mock.DeletePekerjaan(id.Hex())
	if err != nil {
		t.Errorf("Delete error: %v", err)
	}

	if len(mock.Data) != 0 {
		t.Errorf("Expected empty map, got %d items", len(mock.Data))
	}
}
func TestUpdatePekerjaan_NotFound(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	_, err := mock.UpdatePekerjaan(primitive.NewObjectID().Hex(), models.PekerjaanAlumni{})
	if err == nil {
		t.Errorf("Expected error for update non-existing ID, got nil")
	}
}

func TestUpdatePekerjaan_EmptyCompany(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	id := primitive.NewObjectID()
	mock.InsertPekerjaan(models.PekerjaanAlumni{
		ID:             id,
		NamaPerusahaan: "PT Lama",
	})

	_, err := mock.UpdatePekerjaan(id.Hex(), models.PekerjaanAlumni{
		ID:             id,
		NamaPerusahaan: "",
	})

	if err == nil {
		t.Errorf("Expected error for empty NamaPerusahaan")
	}
}
func TestDeletePekerjaan_InvalidID(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	err := mock.DeletePekerjaan("invalid_hex")
	if err == nil {
		t.Errorf("Expected error for invalid ID, got nil")
	}
}
func TestDeletePekerjaan_NotFound(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	err := mock.DeletePekerjaan(primitive.NewObjectID().Hex())
	if err == nil {
		t.Errorf("Expected error for missing ID")
	}
}
func TestFlow_CreateUpdateDelete(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	id := primitive.NewObjectID()

	// CREATE
	mock.InsertPekerjaan(models.PekerjaanAlumni{ID: id, NamaPerusahaan: "PT A", PosisiJabatan: "Dev"})

	// UPDATE
	updated, _ := mock.UpdatePekerjaan(id.Hex(), models.PekerjaanAlumni{
		ID:             id,
		NamaPerusahaan: "PT B",
	})

	if updated.NamaPerusahaan != "PT B" {
		t.Errorf("Expected updated company PT B")
	}

	// DELETE
	_ = mock.DeletePekerjaan(id.Hex())

	if len(mock.Data) != 0 {
		t.Errorf("Expected empty data after delete")
	}
}
func TestCreateMultiple(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	mock.InsertPekerjaan(models.PekerjaanAlumni{ID: primitive.NewObjectID(), NamaPerusahaan: "1"})
	mock.InsertPekerjaan(models.PekerjaanAlumni{ID: primitive.NewObjectID(), NamaPerusahaan: "2"})

	if len(mock.Data) != 2 {
		t.Errorf("Expected 2 data, got %d", len(mock.Data))
	}
}
func TestUpdateWithoutChangingAlumni(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	id := primitive.NewObjectID()
	alumni := primitive.NewObjectID()

	mock.InsertPekerjaan(models.PekerjaanAlumni{
		ID:       id,
		AlumniID: alumni,
	})

	updated, _ := mock.UpdatePekerjaan(id.Hex(), models.PekerjaanAlumni{
		ID:             id,
		AlumniID:       alumni,
		NamaPerusahaan: "New",
	})

	if updated.AlumniID != alumni {
		t.Errorf("AlumniID must not change!")
	}
}
func TestGetPekerjaanByAlumni_Multiple(t *testing.T) {
	mock := repository.NewMockPekerjaanRepo()

	alumni := primitive.NewObjectID()

	mock.InsertPekerjaan(models.PekerjaanAlumni{ID: primitive.NewObjectID(), AlumniID: alumni})
	mock.InsertPekerjaan(models.PekerjaanAlumni{ID: primitive.NewObjectID(), AlumniID: alumni})

	data, _ := mock.GetPekerjaanByAlumni(alumni.Hex())

	if len(data) != 2 {
		t.Errorf("Expected 2 results, got %d", len(data))
	}
}
