package interceptors

import (
	"fmt"
	"log"
	common_mongo "github.com/primitivofr/owly/server/common/mongo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsUserInChatroom allows to check whether the given user belongs to the given chatroom.
// Will return an error if not.
func IsUserInChatroom(chatroomID string, userID string) error {

	userResult, err := common_mongo.FindOneUserCollection(userID)
	if err != nil {
		log.Printf("Error while checking if user belong to chatroom: %v", err)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while checking if user belong to chatroom: %v", err),
		)
	}

	for _, chatroom := range userResult.Chatrooms {
		if chatroom == chatroomID {
			return nil
		}
	}

	log.Printf("User %v not found in this chatroom: %v", userID, chatroomID)
	return status.Errorf(
		codes.Unauthenticated,
		fmt.Sprintf("User %v not found in this chatroom: %v", userID, chatroomID),
	)

}
