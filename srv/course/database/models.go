package database

import "gopkg.in/mgo.v2/bson"

type Course struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Type        string        `bson:"type" json:"type"`
	Topics      []string      `bson:"topics" json:"topics"`
}
