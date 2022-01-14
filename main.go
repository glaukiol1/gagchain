package main

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/glaukiol1/gagchain/blockchain"
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
	tp := blockchain.NewTransactionPool()

	pk, err := crypto.HexToECDSA("d2d9a9aa0fce4a8a1d3141300fa8a0c0087f7ae93dd396d5198381c440584361")
	if err != nil {
		panic(err)
	}
	pb := &pk.PublicKey
	pb1, pk1 := blockchain.Keygen()
	trns1 := bc.NewTransactionInstance(pb, blockchain.PubkeyToAddress(pb1), 5000, tp)
	trns1.Sign(blockchain.PrivateKeyToHex(pk))
	tp.AddTransaction(trns1)
	trns2 := bc.NewTransactionInstance(pb1, blockchain.PubkeyToAddress(pb), 90, tp)
	trns2.Sign(blockchain.PrivateKeyToHex(pk1))

	tp.AddTransaction(trns2)
	bc.AddBlock(tp.MineTransactions())

	// for _, block := range bc.Blocks {
	// 	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	for _, trns := range block.Data {
	// 		fmt.Printf("Transaction From: %s\n", trns.From)
	// 		fmt.Printf("Transaction To: %s\n", trns.To)
	// 		fmt.Printf("Transaction Amount: %s\n", fmt.Sprint(trns.Amount))
	// 		if block.Id != 0 {
	// 			fmt.Printf("Transaction Successfully Signed: %s\n", fmt.Sprint(trns.VerifySignature()))
	// 		}
	// 	}

	// 	pow := blockchain.NewProof(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println("--------------------")
	// }
	dat, _ := json.Marshal(bc.Blocks[1].Data)
	println(string(dat))
	println(bc.GetBalance(blockchain.PubkeyToAddress(pb), tp),
		bc.GetBalance(blockchain.PubkeyToAddress(pb1), tp))
}
