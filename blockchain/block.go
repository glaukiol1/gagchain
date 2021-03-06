package blockchain

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

// TODO: Transactions
// https://dev.to/freakcdev297/creating-transactions-mining-rewards-mint-and-gas-fee-5hhf
type Blockchain struct {
	Blocks []*Block //
}

type Block struct {
	Hash      []byte         //	The block hash
	Data      []*Transaction //	Transactions
	PrevHash  []byte         // Previous Hash, blocks[len(blocks)-1].Hash
	Nonce     int            // Block Nonce (Part of PoW)
	Id        int            // Block ID
	Timestamp int            // timestamp
	Miner     string         // miner address
}

func (bc *Blockchain) AddBlock(data []*Transaction, hash []byte, nonce int, Miner string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	new, err := CreateBlock(data, prevBlock.Hash, prevBlock.Id, bc, hash, nonce, Miner)
	if err != nil {
		log.Fatalln(err)
	}
	bc.Blocks = append(bc.Blocks, new)
}

func CreateBlock(data []*Transaction, prevHash []byte, prevId int, bc *Blockchain, _Hash []byte, _Nonce int, Miner string) (blck *Block, err error) {
	block := &Block{_Hash, data, prevHash, _Nonce, prevId + 1, int(time.Now().Unix()), Miner}

	// check for invalid transactions here
	// make a function to validate transactions
	// and out it of the block
	var invalidBlockError error = errors.New("Invalid Block #" + fmt.Sprint(block.Id))
	if Mining_Node && block.Miner == "" {
		block, err = ValidateBlock(block, bc)
		if err != nil {
			return nil, err
		}
		block.Miner = MyAddress.PublicAddress
		bc.AddRewardTransaction(block)
		pow := NewProof(block)
		nonce, hash := pow.Run()
		block.Hash = hash
		block.Nonce = nonce
		if block.IsValid(bc) {
			return block, nil
		} else {
			return nil, invalidBlockError
		}
	} else {
		block, err = ValidateBlock(block, bc)
		if err != nil {
			return nil, err
		}
		block.Hash = _Hash
		block.Nonce = _Nonce
		pow := NewProof(block)
		if block.Id == 0 {
			nonce, hash := pow.Run()
			block.Nonce = nonce
			block.Hash = hash
		}
		if !pow.Validate() || !block.IsValid(bc) {
			return nil, invalidBlockError
		} else {
			return block, nil
		}
	}

}

func GetGenesis() *Block {
	var to string = "0x4390B0820B4257d8936759e5e043e91a1F9E0BeC"
	public := MintAddress.publicKey
	var GenesisTransaction = &Transaction{int(time.Now().Unix()), []byte(crypto.PubkeyToAddress(*public).Hex()), crypto.FromECDSAPub(public), []byte(to), 100000, []byte{}, [32]byte{}}
	GenesisTransaction.Sign(PrivateKeyToHex(MintAddress.privateKey))
	GenesisTransaction.MakeHash()
	var x []*Transaction
	var b *Blockchain = &Blockchain{}
	block, err := CreateBlock(append(x, GenesisTransaction), []byte{}, -1, b, []byte{}, 113, "")

	if err != nil {
		panic(err)
	}
	return block
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GetGenesis()}}
}
