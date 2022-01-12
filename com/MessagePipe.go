package com

type MessagePipeFunc func(message string)

func MessagePipeStart(port int, cb MessagePipeFunc) {
	StartServer(port, func(message string) {
		cb(message)
	})
}
