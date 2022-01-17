package main

import (
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/glaukiol1/gagchain/blockchain"
	"github.com/glaukiol1/gagchain/com"
	"github.com/glaukiol1/gagchain/db"
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

	// dab.UpdateDB()

	// dat := bc.Blocks[0].Data
	// fmt.Println(string(dat[0].)) // print the genesis block

	// bc := blockchain.InitBlockchain()
	// json.Unmarshal([]byte("{\"Blocks\":[{\"Hash\":\"AACU5KOlFcem3ivACPfUARsOBZttZF3LL5mSfk35B0o=\",\"Data\":[{\"Timestamp\":1642362647,\"From\":\"MHhiNzYyNDg3MjhjZWM5YzM0NEQ4OTA2M0Y4QTYyNjEzOUIwMGM4NTkx\",\"PubKeyBytes\":\"BAJ1WYhU0d1NcKZN1B3pK5arJFJ4RABEFYU/Rb2sGQvbcUxnEabB22/MERFaaXnoe1g/2m0cbMxlZqNWCV3UP48=\",\"To\":\"MHg0MzkwQjA4MjBCNDI1N2Q4OTM2NzU5ZTVlMDQzZTkxYTFGOUUwQmVD\",\"Amount\":100000,\"Signature\":\"uMUGcUJifuBPO1soFnKS3YZ4scoWJW5KcpJiDe7msUNGG8/NoiRJ0xtkjW1hN9EyHKR2B7IG2q5CMxv+knJezAE=\",\"Hash\":[48,56,64,193,166,97,35,89,116,157,197,228,128,18,132,191,59,122,114,59,34,77,40,248,108,231,36,110,41,11,177,66]},{\"Timestamp\":1642362647,\"From\":\"MHhiNzYyNDg3MjhjZWM5YzM0NEQ4OTA2M0Y4QTYyNjEzOUIwMGM4NTkx\",\"PubKeyBytes\":\"BAJ1WYhU0d1NcKZN1B3pK5arJFJ4RABEFYU/Rb2sGQvbcUxnEabB22/MERFaaXnoe1g/2m0cbMxlZqNWCV3UP48=\",\"To\":\"MHg0MzkwQjA4MjBCNDI1N2Q4OTM2NzU5ZTVlMDQzZTkxYTFGOUUwQmVD\",\"Amount\":100,\"Signature\":\"7Byk2z9/pBZ6Bl4D2xJBloNwoWn/kKQrpSARCQwc2JosPnRvLppzPtJA0Mvhzat854bWJlwegFwwuYmkqXsQxAA=\",\"Hash\":[108,191,85,164,44,118,5,174,160,43,141,26,116,16,13,88,214,5,16,238,184,86,15,115,155,169,168,27,162,117,213,148]}],\"PrevHash\":\"\",\"Nonce\":49649,\"Id\":0,\"Timestamp\":1642362647,\"Miner\":\"0x4390B0820B4257d8936759e5e043e91a1F9E0BeC\"},{\"Hash\":\"AAGqnMn3bFT9uuopK1sPp4wfHrrquqfPqqyxciURAdY=\",\"Data\":[{\"Timestamp\":1642285597,\"From\":\"MHg0MzkwQjA4MjBCNDI1N2Q4OTM2NzU5ZTVlMDQzZTkxYTFGOUUwQmVD\",\"PubKeyBytes\":\"BFCmDAdKl+jh1gQhh51gLAsXau2R9pF8t5phzvDRoQ7wVfuDnl8G5KL6kzQbABud6Pydqxzb92L5xPt3yKmIhMk=\",\"To\":\"MHg0MzE1MzFlMTI5MDBCM0IyNjI3YzM3YzMyOWQ2MDJiNzREMDlkM0JC\",\"Amount\":5000,\"Signature\":\"ePysGH1pEEMyknLNsOhhbVfKMTSd2MBgAl/Ir0mKSegJPMihoa+vXXVJP63wfzMTLnwreRR+bHK5Jh3bpNQuRAE=\",\"Hash\":[227,175,196,103,247,165,242,108,65,176,79,223,30,77,235,133,161,59,6,135,74,39,91,217,66,46,3,10,89,51,171,5]},{\"Timestamp\":1642285597,\"From\":\"MHg0MzE1MzFlMTI5MDBCM0IyNjI3YzM3YzMyOWQ2MDJiNzREMDlkM0JC\",\"PubKeyBytes\":\"BIl57r1Hl5cEg7xy1p31TVpDeCevWk8P+l4luTabhwVJoj0zJlnHnzqUmGM/ljbceoosTFGDu92PwXAZGhyjpWU=\",\"To\":\"MHg0MzkwQjA4MjBCNDI1N2Q4OTM2NzU5ZTVlMDQzZTkxYTFGOUUwQmVD\",\"Amount\":90,\"Signature\":\"+3hflJRQJYcmP8vApELFVTuv51noxb0gf3DWSN4mINV8Yy6LDd8DcvUQ377tOvLSBa2zMlH+rV85c9hTX+VaVAA=\",\"Hash\":[73,161,170,213,39,210,244,143,114,23,165,220,48,117,9,100,49,6,32,126,251,139,178,126,126,194,158,3,226,109,2,234]},{\"Timestamp\":1642353671,\"From\":\"MHhiNzYyNDg3MjhjZWM5YzM0NEQ4OTA2M0Y4QTYyNjEzOUIwMGM4NTkx\",\"PubKeyBytes\":\"BAJ1WYhU0d1NcKZN1B3pK5arJFJ4RABEFYU/Rb2sGQvbcUxnEabB22/MERFaaXnoe1g/2m0cbMxlZqNWCV3UP48=\",\"To\":\"MHg0MzkwQjA4MjBCNDI1N2Q4OTM2NzU5ZTVlMDQzZTkxYTFGOUUwQmVD\",\"Amount\":100,\"Signature\":\"mWptPHl7dcxlvMkYQ/WRFi/8h/vR2xzb1z8rurXWYTBkP3UqNF1ONJHueUkkIVlfKSJYaOt5VG9WjawRkenDCgA=\",\"Hash\":[116,27,136,49,190,188,102,191,67,176,178,87,42,234,89,52,17,40,237,73,42,6,37,9,133,45,231,93,139,239,163,18]},{\"Timestamp\":1642362654,\"From\":\"MHhiNzYyNDg3MjhjZWM5YzM0NEQ4OTA2M0Y4QTYyNjEzOUIwMGM4NTkx\",\"PubKeyBytes\":\"BAJ1WYhU0d1NcKZN1B3pK5arJFJ4RABEFYU/Rb2sGQvbcUxnEabB22/MERFaaXnoe1g/2m0cbMxlZqNWCV3UP48=\",\"To\":\"MHg0MzkwQjA4MjBCNDI1N2Q4OTM2NzU5ZTVlMDQzZTkxYTFGOUUwQmVD\",\"Amount\":100,\"Signature\":\"ihalK0cdByOsv6BZ4izopDOTQuAmBkTzRbclvbKvWzIlzXYkEvRlq2+yNoBrMsVq8u5oLpHJXOQiZneBPsYNuwA=\",\"Hash\":[114,216,196,121,239,72,179,245,63,66,64,21,122,64,90,70,82,175,122,192,142,20,139,247,78,199,41,128,94,103,78,247]}],\"PrevHash\":\"AACU5KOlFcem3ivACPfUARsOBZttZF3LL5mSfk35B0o=\",\"Nonce\":22444,\"Id\":1,\"Timestamp\":1642362654,\"Miner\":\"0x4390B0820B4257d8936759e5e043e91a1F9E0BeC\"}]}"), bc)

	tp := blockchain.NewTransactionPool()

	// pb := blockchain.MyAddress.PublicKey
	// pk := blockchain.MyAddress.PrivateKey
	pk1, _ := crypto.HexToECDSA("c6312a1015b5eba3e9707a065f79ff0b35bb6d629e472643e80ed7aa6719c53b") // 0x431531e12900B3B2627c37c329d602b74D09d3BB
	pb1 := &pk1.PublicKey
	// trns1 := bc.NewTransactionInstance(pb, blockchain.PubkeyToAddress(pb1), 5000)
	// trns1.Sign(blockchain.PrivateKeyToHex(pk))
	// tp.AddTransaction(trns1)
	// trns2 := bc.NewTransactionInstance(pb1, blockchain.PubkeyToAddress(pb), 90)
	// trns2.Sign(blockchain.PrivateKeyToHex(pk1))
	// tp.AddTransaction(trns2)
	// bc.AddBlock(tp.MineTransactions(), []byte{}, -1, "")

	// for _, block := range bc.Blocks {
	// 	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("Nonce: %s\n", fmt.Sprint(block.Nonce))
	// 	fmt.Printf("Miner: %s\n", block.Miner)
	// 	fmt.Printf("Valid: %s\n", strconv.FormatBool(block.IsValid(bc)))
	// 	for _, trns := range block.Data {
	// 		fmt.Printf("Transaction From: %s\n", trns.From)
	// 		fmt.Printf("Transaction To: %s\n", trns.To)
	// 		fmt.Printf("Transaction Amount: %s\n", fmt.Sprint(trns.Amount))
	// 		fmt.Printf("Transaction Hash: %x\n", trns.Hash)
	// 		if block.Id != 0 {
	// 			fmt.Printf("Transaction Successfully Signed: %s\n", fmt.Sprint(trns.VerifySignature()))
	// 		}
	// 	}

	// 	pow := blockchain.NewProof(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println("--------------------")
	// }
	// dat, _ := json.Marshal(bc)
	// fmt.Print(string(dat))
	println(bc.GetBalance(blockchain.PubkeyToAddress(blockchain.MyAddress.PublicKey), tp),
		bc.GetBalance(blockchain.PubkeyToAddress(pb1), tp))
	c := make(chan string)
	go func() {
		com.StartHandler()
		c <- "one"
	}()
	com.AddNewNode(":8888")
	time.Sleep(5 * time.Second)
	com.BroadcastMessage(com.MAKE_TYPE_HANDSHAKE("127"))
	for {
		time.Sleep(5 * time.Second)
		println(com.Nodes[0], com.Nodes[1])
	}

}
