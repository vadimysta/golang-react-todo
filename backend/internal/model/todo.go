package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Completed bool `json:"completed" bson:"completed"`
	Body string `json:"body" bson:"body"`
}