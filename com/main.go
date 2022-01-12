package com

import (
	"fmt"
	"strings"
)

func MainTest() {
	MessagePipeStart(8888, func(message Message) {
		fmt.Printf("New message: type: %s data: %s\n", message.msgtype, strings.Join(message.msgdata, " "))
	})
}
