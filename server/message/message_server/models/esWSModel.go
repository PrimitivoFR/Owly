package models

import "primitivofr/owly/server/message/messagepb"

type ESWebSocketResp struct {
	Index     string            `json:"_index"`
	Type      string            `json:"_type"`
	ID        string            `json:"_id"`
	Timestamp string            `json:"_timestamp"`
	Version   int64             `json:"_version"`
	Operation string            `json:"_operation"`
	Source    messagepb.Message `json:"_source"`
}
