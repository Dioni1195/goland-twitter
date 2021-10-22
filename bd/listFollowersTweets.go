package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ListFollowersTweets(ID string, page int) ([]*models.ResponseFollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("relationships")

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}})
	conditions = append(conditions, bson.M{"$skip": (page - 1) * 20})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)
	var result []*models.ResponseFollowersTweets
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return result, false
	}

	return result, true
}
