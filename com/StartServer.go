package com

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func StartServer() {
	// runs on MessagePipeStart(func (message string))
	// used to listen on a port
	// and wait for messages

	serverStart(8888)
}

func handleNewMsg(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := strings.TrimSpace(string(netData))
		println("New Msg: " + msg)
		c.Close()
		break
	}
}

func serverStart(port int) {
	PORT := ":" + fmt.Sprint(port)
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		handleNewMsg(c)
	}
}
