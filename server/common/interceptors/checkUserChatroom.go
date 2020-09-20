package interceptors

import (
	"context"
	"primitivofr/owly/server/common/models"
	common_mongo "primitivofr/owly/server/common/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func IsUserInChatroom(chatroomID string, userID string) (bool, error) {
	filter := bson.M{"_id": userID}
	var user_result models.UserMongo

	err := common_mongo.UserCollection.FindOne(context.Background(), filter).Decode(&user_result)
	if err != nil {
		return false, err
	}

	for _, chatroom := range user_result.Chatrooms {
		if chatroom == chatroomID {
			return true, nil
		}
	}

	return false, nil

}
