package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Username  string             `bson:"username"`
	Chatrooms []string           `bson:"chatrooms"`
}
