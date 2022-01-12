package com

func MainTest() {
	MessagePipeStart(8888, func(message string) {
		println("New message: " + message)
	})
}
