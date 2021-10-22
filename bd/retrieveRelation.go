package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"time"
)

func RetrieveRelation(rel models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("relationships")

	if err := col.FindOne(ctx, rel).Decode(&rel); err != nil {
		return false, err
	}
	return true, nil
}
