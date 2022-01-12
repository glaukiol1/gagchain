package com

import "strings"

type MessagePipeFunc func(message Message)

type MsgType struct {
	name string
}
type Message struct {
	msgtype MsgType
	msgdata []string
}

func MessagePipeStart(port int, cb MessagePipeFunc) {
	StartServer(port, func(message string) {
		cb(Message{MsgType{strings.Split(message, " ")[0]}, strings.Split(message, " ")[1:]})
	})
}
