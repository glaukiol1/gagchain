package main

import (
	"fmt"

	"github.com/glaukiol1/gagchain/blockchain"
	"github.com/glaukiol1/gagchain/db"
)

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
	var bc *blockchain.Blockchain
	if db.DB_DoesExist(db_location) {
		v = dab.ParseDB()
		bc = &blockchain.Blockchain{v} // ignore warning
	} else {
		bc = blockchain.InitBlockchain()
	}

	bc.AddBlock("Block 1")
	dab.Write(bc.RequestChain())

	// dab.UpdateDB()

	fmt.Println(string(bc.Blocks[len(bc.Blocks)-1].Data)) // print the genesis block
}
