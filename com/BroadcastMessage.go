package com

import (
	"bufio"
	"fmt"
	"net"
)

func BroadcastMessage(message string) {
	// connect to the selected node
	// send the message
	// selected node passes the message along.
	sendMessage(message, Nodes)
}

func sendMessage(message string, hosts []string) {
	println(message)
	for _, host := range hosts {
		CONNECT := host
		c, err := net.Dial("tcp", CONNECT)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Write([]byte(message + "\n"))
		bufio.NewReader(c).ReadString('\n')
	}
}
