package blockchain

import "time"

// TODO: Transactions
// https://dev.to/freakcdev297/creating-transactions-mining-rewards-mint-and-gas-fee-5hhf
type Blockchain struct {
	Blocks []*Block //
}

type Block struct {
	Hash      []byte         //
	Data      []*Transaction //
	PrevHash  []byte         //
	Nonce     int            //
	Id        int            //
	Timestamp int            //
}

func (bc *Blockchain) AddBlock(data []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash, prevBlock.Id)
	bc.Blocks = append(bc.Blocks, new)
}

func CreateBlock(data []*Transaction, prevHash []byte, prevId int) *Block {
	block := &Block{[]byte{}, data, prevHash, 0, prevId + 1, int(time.Now().Unix())}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce
	return block
}

func GetGenesis() *Block {
	var GenesisTransaction = NewTransactionInstance("0xtest1", "0xtest2", 10000)
	var x []*Transaction
	return CreateBlock(append(x, GenesisTransaction), []byte{}, -1)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GetGenesis()}}
}
