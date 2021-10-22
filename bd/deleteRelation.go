package bd

import (
	"GitHub/goland-twitter/models"
	"context"
	"time"
)

func DeleteRelation(rel models.Relationship) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("myFirstDatabase")
	col := db.Collection("relationships")

	if _, err := col.DeleteOne(ctx, rel); err != nil {
		return err
	}
	return nil
}
