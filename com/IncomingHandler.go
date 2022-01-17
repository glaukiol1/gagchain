package com

import "encoding/json"

// responders to different input messages

type TYPE_HANDSHAKE struct {
	HandshakeFrom string // handshake IP
}

func StartHandler() {
	MessagePipeStart(8888, func(message Message) {
		switch message.msgtype.name {
		case "TYPE_HANDSHAKE":
			msg := TYPE_HANDSHAKE{}
			json.Unmarshal([]byte(message.msgdata), &msg)
			AddNewNode(msg.HandshakeFrom)
		}
	})
}
