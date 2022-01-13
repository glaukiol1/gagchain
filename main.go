package main

import (
	"encoding/json"
	"fmt"

	"github.com/glaukiol1/gagchain/blockchain"
	"github.com/glaukiol1/gagchain/db"
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
	db_location := "./db/db.db"
	dab := db.GetDB(db_location)
	var v []*blockchain.Block
	json.Unmarshal([]byte(dab.GetContents()), &v)
	var bc *blockchain.Blockchain
	if db.DB_DoesExist(db_location) {
		bc = &blockchain.Blockchain{v} // ignore warning
	} else {
		bc = blockchain.InitBlockchain()
	}

	// dab.Write(bc.RequestChain())

	fmt.Println(string(bc.Blocks[0].Data)) // print the genesis block
}
