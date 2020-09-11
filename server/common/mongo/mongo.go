package common_mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ChatroomCollection *mongo.Collection

func SetupMongoDB() error {
	var err error
	fmt.Println("Connecting to MongoDB...")
	MONGO_HOSTNAME := "localhost"
	MONGO_HOSTNAME = os.Getenv("MONGO_HOSTNAME")
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + MONGO_HOSTNAME + ":27017"))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	//Collections
	ChatroomCollection = client.Database("owly").Collection("chatrooms")
	//

	return nil
}
