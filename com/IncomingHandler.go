package com

import "encoding/json"

// responders to different input messages

type TYPE_HANDSHAKE struct {
	HandshakeFrom string // handshake IP
}

var Nodes []string

func StartHandler() {
	MessagePipeStart(8888, func(message Message) {
		switch message.msgtype.name {
		case "TYPE_HANDSHAKE":
			msg := TYPE_HANDSHAKE{}
			json.Unmarshal([]byte(message.msgdata), &msg)
			Nodes = append(Nodes, msg.HandshakeFrom)
		}
	})
}
