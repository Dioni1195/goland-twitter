package models

type Relationship struct {
	UserID         string `bson:"userid" json:"userId"`
	UserRelationId string `bson:"userrelationid" json:"userRelationId"`
}
