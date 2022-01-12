package com

import "strings"

type MessagePipeFunc func(message Message)

type Message struct {
	msgtype string
	msgdata []string
}

func MessagePipeStart(port int, cb MessagePipeFunc) {
	StartServer(port, func(message string) {
		cb(Message{strings.Split(message, " ")[0], strings.Split(message, " ")[1:]})
	})
}
