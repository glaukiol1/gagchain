package blockchain

import (
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

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
	new := CreateBlock(data, prevBlock.Hash, prevBlock.Id, bc)
	bc.Blocks = append(bc.Blocks, new)
}

func CreateBlock(data []*Transaction, prevHash []byte, prevId int, bc *Blockchain) *Block {
	block := &Block{[]byte{}, data, prevHash, 0, prevId + 1, int(time.Now().Unix())}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	// check for invalid transactions here
	// make a function to validate transactions
	// and out it of the block

	block = ValidateBlock(block, bc)
	block.Hash = hash
	block.Nonce = nonce
	return block
}

func GetGenesis() *Block {
	public, _ := Keygen()
	var to string = "0x4390B0820B4257d8936759e5e043e91a1F9E0BeC" // to address
	var GenesisTransaction = &Transaction{int(time.Now().Unix()), []byte(crypto.PubkeyToAddress(*public).Hex()), crypto.FromECDSAPub(public), []byte(to), 100000, []byte{}, [32]byte{}}
	var x []*Transaction
	var b *Blockchain = &Blockchain{}
	return CreateBlock(append(x, GenesisTransaction), []byte{}, -1, b)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GetGenesis()}}
}
