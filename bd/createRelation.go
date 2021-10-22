package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"time"
)

func CreateRelation(rela models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("relationships")

	_, err := col.InsertOne(ctx, rela)
	if err != nil {
		return false, err
	}
	return true, nil
}
