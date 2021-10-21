package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func CheckUserAlreadyExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("users")

	cond := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, cond).Decode(&result)
	ID := result.Id.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
