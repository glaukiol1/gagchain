package com

import "encoding/json"

// message builder
// build messages from templates
// like transactions
// or request the chain
// to be ready to broadcast
//TODO-<

func MAKE_TYPE_HANDSHAKE(host string) string {
	d, err := json.Marshal(TYPE_HANDSHAKE{host})
	if err != nil {
		panic(err)
	}
	return "TYPE_HANDSHAKE " + string(d)
}
