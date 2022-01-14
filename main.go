package main

import (
	"fmt"

	"github.com/glaukiol1/gagchain/blockchain" // blockchain main module
	// custom database
)

func main() {
	// chain := blockchain.InitBlockchain()
	// trns1 := blockchain.NewTransactionInstance("0xtest1", "0xtest2", 10000)
	// trns2 := blockchain.NewTransactionInstance("0xtest2", "0xtest3", 5000)
	// var arr1 []*blockchain.Transaction
	// chain.AddBlock(append(arr1, trns1, trns2))

	// var arr2 []*blockchain.Transaction
	// chain.AddBlock(append(arr2, trns2))
	// for _, block := range chain.Blocks {
	// 	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Hash: %x\n", block.Hash)

	// 	for _, trns := range block.Data {
	// 		fmt.Printf("Transaction From: %s\n", trns.From)
	// 		fmt.Printf("Transaction To: %s\n", trns.To)
	// 		fmt.Printf("Transaction Amount: %s\n", fmt.Sprint(trns.Amount))
	// 	}

	// 	pow := blockchain.NewProof(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println("--------------------")
	// }
	// println("0xtest1 Balance: " + fmt.Sprint(chain.GetBalance("0xtest1")))
	// println("0xtest2 Balance: " + fmt.Sprint(chain.GetBalance("0xtest2")))
	// println("0xtest3 Balance: " + fmt.Sprint(chain.GetBalance("0xtest3")))

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

	trns2 := blockchain.NewTransactionInstance("0xtest2", "0xtest3", 5000)
	pb, pk := blockchain.Keygen()
	// pb1, _ := blockchain.Keygen()
	fmt.Println(blockchain.PubkeyToAddress(pb))
	trns2.Sign(blockchain.PrivateKeyToHex(pk))
}
