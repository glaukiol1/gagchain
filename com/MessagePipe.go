package com

type MessagePipeFunc func(message string)

func MessagePipeStart(cb MessagePipeFunc) {
	// listen on a port
	// wait for connections
	// on a connection, call the `cb`
	// with the message as a parameter
}
