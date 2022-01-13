package com

import "fmt"

// message builder
// build messages from templates
// like transactions
// or request the chain
// to be ready to broadcast
//TODO-<

func TYPE_GET_CHAIN() string {
	return "TYPE_GET_CHAIN full_chain"
}

func TYPE_SEND_TRANSACTION() // todo

func TYPE_GET_BLOCKS_SINCE(since int) string {
	return "TYPE_GET_BLOCKS_SINCE " + fmt.Sprint(since)
}
