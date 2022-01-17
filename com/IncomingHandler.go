package com

import (
	"encoding/json"
	"time"
)

// responders to different input messages

type TYPE_HANDSHAKE struct {
	HandshakeFrom string // handshake IP
}

func StartHandler() {
	AddNewNode(":8888")
	go func() {
		startHandler()
	}()
	time.Sleep(1 * time.Second)
	return
}

func startHandler() {
	MessagePipeStart(8888, func(message Message) {
		switch message.msgtype.name {
		case "TYPE_HANDSHAKE":
			msg := TYPE_HANDSHAKE{}
			json.Unmarshal([]byte(message.msgdata), &msg)
			AddNewNode(msg.HandshakeFrom)
		}
	})
}
