package main

import (
	"fmt"
	"strconv"

	"github.com/glaukiol1/gagchain/blockchain" // blockchain main module
	// custom database
)

func main() {
	// trns1 := blockchain.NewTransactionInstance("0xtest1", "0xtest2", 10000)
	// trns2 := blockchain.NewTransactionInstance("0xtest2", "0xtest3", 5000)
	// var arr1 []*blockchain.Transaction
	// chain.AddBlock(append(arr1, trns1, trns2))

	// var arr2 []*blockchain.Transaction
	// chain.AddBlock(append(arr2, trns2))

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

	bc := blockchain.InitBlockchain()

	pb, pk := blockchain.Keygen()
	pb1, pk1 := blockchain.Keygen()
	trns1 := blockchain.NewTransactionInstance(pb, "0xtest3", 5000)
	trns2 := blockchain.NewTransactionInstance(pb1, "0xtest3", 5000)
	trns1.Sign(blockchain.PrivateKeyToHex(pk))
	trns2.Sign(blockchain.PrivateKeyToHex(pk1))

	tp := blockchain.NewTransactionPool()
	tp.AddTransaction(trns1)
	tp.AddTransaction(trns2)
	bc.AddBlock(tp.MineTransactions())

	for _, block := range bc.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		for _, trns := range block.Data {
			fmt.Printf("Transaction From: %s\n", trns.From)
			fmt.Printf("Transaction To: %s\n", trns.To)
			fmt.Printf("Transaction Amount: %s\n", fmt.Sprint(trns.Amount))
			if block.Id != 0 {
				fmt.Printf("Transaction Successfully Signed: %s\n", fmt.Sprint(trns.VerifySignature()))
			}
		}

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("--------------------")
	}
}
