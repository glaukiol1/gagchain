package blockchain

type Blockchain struct {
	Blocks []*Block //
}

type Block struct {
	Hash     []byte //
	Data     []byte //
	PrevHash []byte //
	Nonce    int    //
}

func (bc *Blockchain) AddBlock(data string) {
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash
	new := CreateBlock(data, prevHash)
	bc.Blocks = append(bc.Blocks, new)
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce
	return block
}

func GetGenesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GetGenesis()}}
}
