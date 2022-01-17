package com

import (
	"fmt"
)

func AddNewNode(node string) {
	// maybe some checks about the node

	// todo
	if !HasNode(node) {
		Nodes = append(Nodes, node)
	} else {
		// already has node
		fmt.Println("Already have node")
	}
}

func HasNode(node string) bool {
	for _, n := range Nodes {
		if n == node {
			return true
		}
	}
	return false
}
