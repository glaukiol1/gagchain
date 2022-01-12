package com

import (
	"fmt"
	"strings"
)

func MainTest(port int) {
	MessagePipeStart(port, func(message Message) {
		fmt.Printf("New message: type: %s data: %s\n", message.msgtype, strings.Join(message.msgdata, " "))
	})
}

func MessageSendTest(selctedNode string) {
	BroadcastMessage("TYPE_NEW_TRANSACTION new_transaction_data reciver_data", selctedNode)
}
