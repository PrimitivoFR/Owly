package models

import "primitivofr/owly/server/message/messagepb"

// ESWebSocketResp is a structure that matches the webSocket messages sent by the ES plugin
// Every time something happens, we get notified.
// Regarding operation, several values are possible :
//  "CREATE" => Document has been created
//  "INDEX" => Document has been updated
//  "DELETE" => Document has been deleted |
type ESWebSocketResp struct {
	Index     string            `json:"_index"`
	Type      string            `json:"_type"`
	ID        string            `json:"_id"`
	Timestamp string            `json:"_timestamp"`
	Version   int64             `json:"_version"`
	Operation string            `json:"_operation"`
	Source    messagepb.Message `json:"_source"`
}
