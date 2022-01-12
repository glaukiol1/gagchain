package com

func StartServer() {
	// runs on MessagePipeStart(func (message string))
	// used to listen on a port
	// and wait for messages
}

func BroadcastMessage(message string) {
	// connect to the selected node
	// send the message
	// selected node passes the message along.
}

type MessagePipeFunc func(message string)

func MessagePipeStart(cb MessagePipeFunc) {
	// listen on a port
	// wait for connections
	// on a connection, call the `cb`
	// with the message as a parameter
}
