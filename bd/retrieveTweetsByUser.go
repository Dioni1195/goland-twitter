package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func RetrieveTweetsByUser(userId string, page int64) (ts []models.RetrieveTweet, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("tweets")

	cond := bson.M{"userid": userId}
	opts := options2.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "date", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, cond, opts)
	if err != nil {
		log.Fatalln("Error retriving tweets by user " + err.Error())
		return
	}

	if err = cursor.All(context.TODO(), &ts); err != nil {
		return
	}

	return
}
