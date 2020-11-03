package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chatroom struct {
	ID    primitive.ObjectID `bson:"_id, omitempty"`
	Owner string             `bson:"owner,omitempty`
	Name  string             `bson:"name"`
	Users []string           `bson:"users"`
}
