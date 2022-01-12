package main

import (
	"github.com/glaukiol1/gagchain/com"
)

// TODO: Make a TCP server for communication between nodes.

func main() {
	// chain := blockchain.InitBlockchain()
	// chain.AddBlock("Block 1 jkanjfbhfbofuoiy41y3o12y41ou4y1u4yuOYUO$YOUY!OUIY$@UIO!Y$!OI$@!")
	// chain.AddBlock("Block 2")
	// chain.AddBlock("Block 3")
	// for _, block := range chain.Blocks {
	// 	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("Data: %s\n", block.Data)

	// 	pow := blockchain.NewProof(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println("--------------------")
	// }

	com.MainTest()
}
