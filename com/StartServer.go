package com

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type fn func(message string)

func StartServer(port int, cb fn) {
	// runs on MessagePipeStart(func (message string))
	// used to listen on a port
	// and wait for messages

	serverStart(port, func(message string) {
		cb(message)
	})
}

func handleNewMsg(c net.Conn) string {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return "ERROR"
		}
		msg := strings.TrimSpace(string(netData))
		c.Close()
		return msg
	}
}

type serverStartFunc func(message string)

func serverStart(port int, cb serverStartFunc) {
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
		cb(handleNewMsg(c))
	}
}
