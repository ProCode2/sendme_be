package main

import (
	"encoding/json"
	"log"
)

const (
	SENDMESSAGEACTION      = "send-message"
	JOINROOMACTION         = "join-room"
	LEAVEROOMACTION        = "leave-room"
	NOTIFYOTHERUSERJOINING = "other-user-joined"
)

type Message struct {
	Action  string  `json:"action"`
	Message string  `json:"message"`
	Target  string  `json:"target"`
	Sender  *Client `json:"sender"`
}

func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return json
}
