package repository

import (
	"crud-alumni/app/models"
	"crud-alumni/database"
	"go.mongodb.org/mongo-driver/bson"
)


// Ambil user berdasarkan email
func FindUserByEmail(email string) (*models.User, error) {
	collection := database.DB.Collection("users")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
