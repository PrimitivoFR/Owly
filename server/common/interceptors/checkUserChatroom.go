package interceptors

import (
	common_mongo "primitivofr/owly/server/common/mongo"
)

func IsUserInChatroom(chatroomID string, userID string) (bool, error) {

	userResult, err := common_mongo.FindOneUserCollection(userID)
	if err != nil {
		return false, err
	}

	for _, chatroom := range userResult.Chatrooms {
		if chatroom == chatroomID {
			return true, nil
		}
	}

	return false, nil

}
