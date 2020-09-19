package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chatroom struct {
	ID    primitive.ObjectID `bson:"_id, omitempty"`
	Name  string             `bson:"name"`
	Users []string           `bson:"users"`
}
