package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func UpdateUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("users")

	record := make(map[string]interface{})
	if len(u.Name) > 0 {
		record["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		record["lastName"] = u.LastName
	}
	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}
	if len(u.Bio) > 0 {
		record["bio"] = u.Bio
	}
	if len(u.Location) > 0 {
		record["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		record["webSite"] = u.WebSite
	}
	record["birthDate"] = u.BirthDate

	recordToUpdate := bson.M{
		"$set": record,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}

	_, error := col.UpdateOne(ctx, filter, recordToUpdate)
	if error != nil {
		return false, error
	}

	return true, nil
}
