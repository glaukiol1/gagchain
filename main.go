package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block //
}

type Block struct {
	Hash     []byte //
	Data     []byte //
	PrevHash []byte //
}

func (b *Block) AssignHash() {
	res := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(res)
	b.Hash = hash[:]
}

func (bc *Blockchain) AddBlock(data string) {
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash
	new := CreateBlock(data, prevHash)
	bc.Blocks = append(bc.Blocks, new)
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.AssignHash()
	return block
}

func GetGenesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GetGenesis()}}
}

func main() {
	chain := InitBlockchain()
	chain.AddBlock("Block 1")
	chain.AddBlock("Block 2")
	chain.AddBlock("Block 3")
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println("--------------------")
	}
}
