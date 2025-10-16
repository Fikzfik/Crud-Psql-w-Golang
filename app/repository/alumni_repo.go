package repository

import (
	"crud-alumni/app/models"
	"crud-alumni/database"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// Ambil semua alumni berdasarkan fakultas
func GetAllAlumniByFak(fak string) ([]models.Alumni, error) {
	collection := database.DB.Collection("alumni")

	filter := bson.M{"fakultas": fak}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.Alumni
	for cursor.Next(ctx) {
		var a models.Alumni
		if err := cursor.Decode(&a); err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

// Ambil semua alumni
func GetAllAlumni() ([]models.Alumni, error) {
	collection := database.DB.Collection("alumni")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.Alumni
	for cursor.Next(ctx) {
		var a models.Alumni
		if err := cursor.Decode(&a); err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

// Ambil alumni berdasarkan ID
func GetAlumniByID(id string) (models.Alumni, error) {
	collection := database.DB.Collection("alumni")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Alumni{}, fmt.Errorf("invalid ObjectID: %v", err)
	}

	var a models.Alumni
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&a)
	return a, err
}

// Tambah alumni baru
func InsertAlumni(a models.Alumni) error {
	collection := database.DB.Collection("alumni")

	a.ID = primitive.NewObjectID()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, a)
	return err
}

// Update alumni
func UpdateAlumni(id string, a models.Alumni) error {
	collection := database.DB.Collection("alumni")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"nim":          a.NIM,
			"nama":         a.Nama,
			"jurusan":      a.Jurusan,
			"angkatan":     a.Angkatan,
			"tahun_lulus":  a.TahunLulus,
			"no_telepon":   a.NoTelepon,
			"alamat":       a.Alamat,
			"fakultas":     a.Fakultas,
			"updated_at":   time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

// Hapus alumni
func DeleteAlumni(id string) error {
	collection := database.DB.Collection("alumni")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %v", err)
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// Ambil alumni dengan pagination + search + sort
func GetAlumniWithPagination(search, sortBy, order string, limit, page int) ([]models.Alumni, error) {
	collection := database.DB.Collection("alumni")

	// Filtering
	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"nama": bson.M{"$regex": search, "$options": "i"}},
				{"jurusan": bson.M{"$regex": search, "$options": "i"}},
				{"fakultas": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	// Sorting
	sortOrder := 1
	if order == "desc" {
		sortOrder = -1
	}
	findOptions := options.Find().
		SetSort(bson.D{{Key: sortBy, Value: sortOrder}}).
		SetLimit(int64(limit)).
		SetSkip(int64((page - 1) * limit))

	// Query
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.Alumni
	for cursor.Next(ctx) {
		var a models.Alumni
		if err := cursor.Decode(&a); err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

// Hitung total alumni (untuk pagination)
func CountAlumni(search string) (int64, error) {
	collection := database.DB.Collection("alumni")

	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"nama": bson.M{"$regex": search, "$options": "i"}},
				{"jurusan": bson.M{"$regex": search, "$options": "i"}},
				{"fakultas": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	count, err := collection.CountDocuments(ctx, filter)
	return count, err
}
	