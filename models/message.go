package models

import (
	"encoding/json"
	"time"
)

type MessageType string

const (
	TypeKeyExchange MessageType = "key_keyexchange"
	TypeChat        MessageType = "chat"
	TypeUserList    MessageType = "user_list"
	TypeUserJoined  MessageType = "user_joined"
	TypeUserLeft    MessageType = "user_left"
	TypeError       MessageType = "error"
)

type Message struct {
	Type      MessageType `json:"type"`
	From      string      `json:"from,omitempty"`
	To        string      `json:"to,omitempty"`
	Content   string      `json:"content,omitempty"`
	PublicKey string      `json:"public_key,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	Encrypted bool        `json:"encrypted,omitempty"`
}

func (m *Message) ToJSON() []byte {
	data, _ := json.Marshal(m)
	return data
}
func MessageFromJSON(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return &msg, err
}
