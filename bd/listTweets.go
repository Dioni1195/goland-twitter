package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ListTweets() ([]models.RetrieveTweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("tweets")

	var tweets []models.RetrieveTweet

	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return tweets, err
	}

	if err = cursor.All(ctx, &tweets); err != nil {
		return tweets, err
	}

	return tweets, nil
}
