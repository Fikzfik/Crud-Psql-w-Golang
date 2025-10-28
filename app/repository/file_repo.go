package repository

import (
	"context"
	"crud-alumni/app/models"
	"crud-alumni/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ===== Foto Repository =====

func InsertFoto(f models.File) error {
	println("test")
	collection := database.DB.Collection("foto")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	f.ID = primitive.NewObjectID()
	f.UploadedAt = time.Now()

	_, err := collection.InsertOne(ctx, f)
	return err
}

func GetAllFoto() ([]models.File, error) {
	collection := database.DB.Collection("foto")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
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
	collection := database.DB.Collection("foto")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.File{}, err
	}

	var f models.File
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&f)
	return f, err
}

func DeleteFoto(id string) error {
	collection := database.DB.Collection("foto")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// ===== Sertifikat Repository =====

func InsertSertifikat(s models.File) error {
	collection := database.DB.Collection("sertifikat")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.ID = primitive.NewObjectID()
	s.UploadedAt = time.Now()

	_, err := collection.InsertOne(ctx, s)
	return err
}

func GetAllSertifikat() ([]models.File, error) {
	collection := database.DB.Collection("sertifikat")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
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
	collection := database.DB.Collection("sertifikat")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.File{}, err
	}

	var s models.File
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&s)
	return s, err
}

func DeleteSertifikat(id string) error {
	collection := database.DB.Collection("sertifikat")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
