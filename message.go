package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

const (
	SENDMESSAGEACTION      = "send-message"
	JOINROOMACTION         = "join-room"
	LEAVEROOMACTION        = "leave-room"
	NOTIFYOTHERUSERJOINING = "other-user-joined"

	CREATEROOMANDSENDID = "create-room-and-send-id"
	CREATEDROOM         = "created-room"
)

type MessageData struct {
	Id          uuid.UUID `json:"id"`
	MessageType string    `json:"type"`
	Filename    string    `json:"filename"`
	Data        string    `json:"data"`
	Size        int       `json:"size"`
}

type Message struct {
	Action  string       `json:"action"`
	Message *MessageData `json:"message"`
	Target  string       `json:"target"`
	Sender  *Client      `json:"sender"`
}

func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return json
}
