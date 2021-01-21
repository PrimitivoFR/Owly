package common_mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Collections: Add your collection global var here
var ChatroomCollection *mongo.Collection
var UserCollection *mongo.Collection

//
func SetupMongoDB() error {
	var err error
	fmt.Println("Connecting to MongoDB...")
	// MONGO_HOSTNAME := "localhost"
	MONGO_HOSTNAME := os.Getenv("MONGO_HOSTNAME")
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + MONGO_HOSTNAME + ":27017"))
	if err != nil {
		return err
	}
	log.Println("Connected to mongo !")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	//Collections: Fill you collection global var here
	ChatroomCollection = client.Database("owly").Collection("chatrooms")
	UserCollection = client.Database("owly").Collection("users")
	//

	return nil
}

//
