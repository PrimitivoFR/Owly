package models

type User struct {
	// ID        primitive.ObjectID `bson:"_id, omitempty"`
	ID        string   `bson:"_id, omitempty"`
	Username  string   `bson:"username"`
	Chatrooms []string `bson:"chatrooms"`
}
