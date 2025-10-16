package repository

import (
	"context"
	"crud-alumni/app/models"
	"crud-alumni/database"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

// Ambil semua pekerjaan
func GetAllPekerjaan() ([]models.PekerjaanAlumni, error) {
	collection := database.DB.Collection("pekerjaan_alumni")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.PekerjaanAlumni
	for cursor.Next(ctx) {
		var p models.PekerjaanAlumni
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// Ambil pekerjaan berdasarkan ID
func GetPekerjaanByID(id string) (models.PekerjaanAlumni, error) {
	collection := database.DB.Collection("pekerjaan_alumni")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.PekerjaanAlumni{}, fmt.Errorf("invalid ObjectID: %v", err)
	}

	var p models.PekerjaanAlumni
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&p)
	return p, err
}

// Ambil pekerjaan berdasarkan alumni
func GetPekerjaanByAlumni(alumniID string) ([]models.PekerjaanAlumni, error) {
	collection := database.DB.Collection("pekerjaan_alumni")

	alumniObjID, err := primitive.ObjectIDFromHex(alumniID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID: %v", err)
	}

	cursor, err := collection.Find(ctx, bson.M{"alumni_id": alumniObjID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.PekerjaanAlumni
	for cursor.Next(ctx) {
		var p models.PekerjaanAlumni
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// Tambah pekerjaan baru
func InsertPekerjaan(p models.PekerjaanAlumni) error {
	collection := database.DB.Collection("pekerjaan_alumni")

	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.IsDeleted = false

	_, err := collection.InsertOne(ctx, p)
	return err
}

// Update pekerjaan
func UpdatePekerjaan(id string, p models.PekerjaanAlumni) error {
	collection := database.DB.Collection("pekerjaan_alumni")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"nama_perusahaan":      p.NamaPerusahaan,
			"posisi_jabatan":       p.PosisiJabatan,
			"bidang_industri":      p.BidangIndustri,
			"lokasi_kerja":         p.LokasiKerja,
			"gaji_range":           p.GajiRange,
			"tanggal_mulai_kerja":  p.TanggalMulaiKerja,
			"tanggal_selesai_kerja": p.TanggalSelesaiKerja,
			"status_pekerjaan":     p.StatusPekerjaan,
			"deskripsi_pekerjaan":  p.DeskripsiPekerjaan,
			"updated_at":           time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

// Soft delete pekerjaan
func SoftDeletePekerjaan(id string, isDelete bool) error {
	collection := database.DB.Collection("pekerjaan_alumni")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %v", err)
	}

	update := bson.M{"$set": bson.M{"isdelete": isDelete, "updated_at": time.Now()}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

// Hapus permanen pekerjaan
func DeletePekerjaan(id string) error {
	collection := database.DB.Collection("pekerjaan_alumni")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %v", err)
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// Ambil pekerjaan aktif (isdelete = false)
func GetPekerjaanAktif() ([]models.PekerjaanAlumni, error) {
	collection := database.DB.Collection("pekerjaan_alumni")

	cursor, err := collection.Find(ctx, bson.M{"isdelete": false})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.PekerjaanAlumni
	for cursor.Next(ctx) {
		var p models.PekerjaanAlumni
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// Ambil dengan pagination
func GetPekerjaanWithPagination(search, sortBy, order string, limit, page int) ([]models.PekerjaanAlumni, error) {
	collection := database.DB.Collection("pekerjaan_alumni")

	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"nama_perusahaan": bson.M{"$regex": search, "$options": "i"}},
				{"posisi_jabatan": bson.M{"$regex": search, "$options": "i"}},
				{"bidang_industri": bson.M{"$regex": search, "$options": "i"}},
				{"lokasi_kerja": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	sortOrder := 1
	if order == "desc" {
		sortOrder = -1
	}

	opts := options.Find().
		SetSort(bson.D{{Key: sortBy, Value: sortOrder}}).
		SetLimit(int64(limit)).
		SetSkip(int64((page - 1) * limit))

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.PekerjaanAlumni
	for cursor.Next(ctx) {
		var p models.PekerjaanAlumni
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// Hitung total pekerjaan (untuk pagination)
func CountPekerjaan(search string) (int64, error) {
	collection := database.DB.Collection("pekerjaan_alumni")

	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"nama_perusahaan": bson.M{"$regex": search, "$options": "i"}},
				{"posisi_jabatan": bson.M{"$regex": search, "$options": "i"}},
				{"bidang_industri": bson.M{"$regex": search, "$options": "i"}},
				{"lokasi_kerja": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	count, err := collection.CountDocuments(ctx, filter)
	return count, err
}
