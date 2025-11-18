package repository

import (
	"context"
	"crud-alumni/app/models"
	"crud-alumni/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ===== Repository Gabungan (Collection: files) =====

// InsertFoto - menyimpan data foto ke collection files
func InsertFoto(f models.File) error {
	collection := database.DB.Collection("files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	f.ID = primitive.NewObjectID()
	f.UploadedAt = time.Now()

	_, err := collection.InsertOne(ctx, f)
	return err
}

// GetAllFoto - mengambil semua data file yang bertipe foto
func GetAllFoto() ([]models.File, error) {
	collection := database.DB.Collection("files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// hanya ambil file bertipe image
	filter := bson.M{"file_type": bson.M{"$regex": "^image/"}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.File
	for cursor.Next(ctx) {
		var f models.File
		if err := cursor.Decode(&f); err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}

func GetFotoByID(id string) (models.File, error) {
	collection := database.DB.Collection("files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.File{}, err
	}

	var f models.File
	err = collection.FindOne(ctx, bson.M{"_id": objID, "file_type": bson.M{"$regex": "^image/"}}).Decode(&f)
	return f, err
}

func DeleteFoto(id string) error {
	collection := database.DB.Collection("files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID, "file_type": bson.M{"$regex": "^image/"}})
	return err
}

// ===== Sertifikat (tetap pakai fungsi terpisah tapi 1 collection) =====

func InsertSertifikat(s models.File) error {
	collection := database.DB.Collection("files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.ID = primitive.NewObjectID()
	s.UploadedAt = time.Now()

	_, err := collection.InsertOne(ctx, s)
	return err
}

func GetAllSertifikat() ([]models.File, error) {
	collection := database.DB.Collection("files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// hanya ambil file bertipe PDF
	filter := bson.M{"file_type": "application/pdf"}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.File
	for cursor.Next(ctx) {
		var s models.File
		if err := cursor.Decode(&s); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

func GetSertifikatByID(id string) (models.File, error) {
	collection := database.DB.Collection("files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.File{}, err
	}

	var s models.File
	err = collection.FindOne(ctx, bson.M{"_id": objID, "file_type": "application/pdf"}).Decode(&s)
	return s, err
}

func DeleteSertifikat(id string) error {
	collection := database.DB.Collection("files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID, "file_type": "application/pdf"})
	return err
}
