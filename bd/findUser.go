package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func FindUser(ID string) (user models.User, error error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("users")

	objID, _ := primitive.ObjectIDFromHex(ID)

	cond := bson.M{"_id": objID}

	err := col.FindOne(ctx, cond).Decode(&user)
	user.Password = ""
	if err != nil {
		log.Fatalln("Not found" + err.Error())
		return user, err
	}

	return user, nil
}
