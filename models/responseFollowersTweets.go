package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ResponseFollowersTweets struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userId"`
	UserRelationId string             `bson:"userrelationid" json:"userRelationId"`
	Tweet          struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Message string    `bson: "message" json:"message,omitempty"`
		Date    time.Time `bson: "date" json:"date,omitempty"`
	}
}
