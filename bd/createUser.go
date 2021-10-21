package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("users")

	u.Password, _ = CryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil

}
