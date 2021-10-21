package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func DeleteTweet(ID string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)
	cond := bson.M{
		"_id":    objID,
		"userid": userID,
	}

	_, err := col.DeleteOne(ctx, cond)
	return err

}
