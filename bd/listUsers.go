package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ListUsers(ID string, page int64, search string, param string) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("users")

	var users []*models.User

	findOpts := options2.Find()
	findOpts.SetSkip((page - 1) * 20)
	findOpts.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOpts)
	if err != nil {
		log.Print("List users Error: " + err.Error())
		return users, false
	}

	var found, include bool

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Print("Error decoding record: " + err.Error())
			return users, false
		}

		rel := models.Relationship{
			UserID:         ID,
			UserRelationId: user.Id.Hex(),
		}

		include = false

		found, err = RetrieveRelation(rel)
		if param == "new" && !found {
			include = true
		}
		if param == "follow" && found {
			include = true
		}

		if rel.UserRelationId == ID {
			include = false
		}

		if include {
			user.Password = ""
			user.Bio = ""
			user.Location = ""
			user.Avatar = ""
			user.Banner = ""

			users = append(users, &user)
		}

	}

	err = cursor.Err()
	if err != nil {
		log.Print("Cursor users Error: " + err.Error())
		return users, false
	}

	cursor.Close(context.TODO())
	return users, true
}
