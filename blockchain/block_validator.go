package blockchain

import (
	"errors"
	"fmt"
)

func ValidateBlock(block *Block, bc *Blockchain) (*Block, error) {
	var twoMintAddressTransactionsError = errors.New("Two transactions from MintAddress in the same block")
	if block.Id == 0 { // if the block is the genesis block
		return block, nil
	} else {
		oldTransactions := block.Data
		var newTransactions []*Transaction // these are verified
		newBlock := &Block{[]byte{}, newTransactions, block.PrevHash, -1, block.Id, block.Timestamp, block.Miner}
		hasMintAddressTransaction := false
		for _, oldTrns := range oldTransactions {
			if string(oldTrns.From) == PubkeyToAddress(MintAddress.publicKey) {
				newTransactions = append(newTransactions, oldTrns)
				if hasMintAddressTransaction {
					return nil, twoMintAddressTransactionsError
				}
				hasMintAddressTransaction = true
			} else {
				if oldTrns.IsValid() && oldTrns.IsReady() && !(bc.GetBalance(string(oldTrns.From), &TransactionPool{newTransactions /* this is safe, since {newTransaction} is only verified transactions */}) < oldTrns.Amount) && oldTrns.VerifySignature() {
					newTransactions = append(newTransactions, oldTrns)
				} else {
					println("Invalid transaction")
				}
			}

		}
		newBlock.Data = newTransactions
		return newBlock, nil
	}
}

func (bc *Blockchain) AddRewardTransaction(block *Block) {
	rewardTransaction := bc.NewTransactionInstance(MintAddress.publicKey, block.Miner, Reward)
	rewardTransaction.Sign(PrivateKeyToHex(MintAddress.privateKey))
	rewardTransaction.MakeHash()
	block.Data = append(block.Data, rewardTransaction)
}

func (block *Block) IsValid(bc *Blockchain) bool {
	if block.Id == 0 {
		return true // change this, security issue
	}
	if !(len(fmt.Sprintf("%x", block.Hash)) == 64) {
		println("Hash is wrong size")
		return false
	} else {
		if block.Miner == "" {
			println("Miner is blank")
			return false
		} else {
			pow := NewProof(block)
			if !pow.Validate() {
				println("PoW validation failed")
				return false
			} else {
				if block.Timestamp < (bc.Blocks[block.Id-1].Timestamp) {
					println("Block timestamp is incorrect")
					return false
				} else {
					if block.Id != len(bc.Blocks[:block.Id]) {
						println("Block ID is wrong ", len(bc.Blocks[:block.Id]), block.Id)
						return false
					} else {
						if string(block.PrevHash) != string(bc.Blocks[block.Id-1].Hash) {
							println("Block PrevHash is wrong")
							return false
						} else {
							if block.Miner != string(block.Data[len(block.Data)-1].To) {
								println("Block Miner is incorrect")
								return false
							} else {
								hasMintTransaction := false
								for _, trns := range block.Data {
									if string(trns.From) == PubkeyToAddress(MintAddress.publicKey) {
										if hasMintTransaction {
											println("More than one MintAddress transaction")
											return false
										}
										hasMintTransaction = true
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return true
}
