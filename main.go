package main

import (
	"fmt"
	"strconv"

	"github.com/glaukiol1/gagchain/blockchain" // blockchain main module
	// custom database
)

func main() {
	chain := blockchain.InitBlockchain()
	trns1 := blockchain.NewTransactionInstance([]byte("from:MINER_ADDRESS"), []byte("to:ad1"))
	trns2 := blockchain.NewTransactionInstance([]byte("from:MINER_ADDRESS"), []byte("to:ad2"))
	var arr1 []*blockchain.Transaction
	chain.AddBlock(append(arr1, trns1, trns2))

	// var arr2 []*blockchain.Transaction
	// chain.AddBlock(append(arr2, trns2))
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)

		for _, trns := range block.Data {
			fmt.Printf("Transaction Input: %s\n", trns.Input)
			fmt.Printf("Transaction Output: %s\n", trns.Output)
		}

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("--------------------")
	}

	// db_location := "./db/db.db"
	// dab := db.GetDB(db_location)
	// var v []*blockchain.Block
	// var bc *blockchain.Blockchain
	// if db.DB_DoesExist(db_location) {
	// 	v = dab.ParseDB()
	// 	bc = &blockchain.Blockchain{v} // ignore warning
	// } else {
	// 	bc = blockchain.InitBlockchain()
	// }

	// // dab.UpdateDB()

	// dat := bc.Blocks[0].Data
	// fmt.Println(string(dat[0].Output)) // print the genesis block
}
