package blockchain

import "time"

type Blockchain struct {
	Blocks []*Block //
}

type Block struct {
	Hash      []byte //
	Data      []byte //
	PrevHash  []byte //
	Nonce     int    //
	Id        int    //
	Timestamp int    //
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash, prevBlock.Id)
	bc.Blocks = append(bc.Blocks, new)
}

func CreateBlock(data string, prevHash []byte, prevId int) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0, prevId + 1, int(time.Now().Unix())}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce
	return block
}

func GetGenesis() *Block {
	return CreateBlock("Genesis", []byte{}, -1)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GetGenesis()}}
}
