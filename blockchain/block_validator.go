package blockchain

func ValidateBlock(block *Block, bc *Blockchain) *Block {
	if block.Id == 0 { // if the block is the genesis block
		return block
	} else {
		oldTransactions := block.Data
		var newTransactions []*Transaction // these are verified
		newBlock := &Block{[]byte{}, newTransactions, block.PrevHash, -1, block.Id, block.Timestamp, block.Miner}
		hasMintAddressTransaction := false
		for _, oldTrns := range oldTransactions {
			if string(oldTrns.From) == PubkeyToAddress(MintAddress.publicKey) {
				println("found reward transaction.")
				newTransactions = append(newTransactions, oldTrns)
				if hasMintAddressTransaction {
					panic("Found two transactions from MintAddress in the same blocl")
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
		return newBlock
	}
}

func (bc *Blockchain) AddRewardTransaction(block *Block) {
	rewardTransaction := bc.NewTransactionInstance(MintAddress.publicKey, block.Miner, Reward)
	rewardTransaction.Sign(PrivateKeyToHex(MintAddress.privateKey))
	rewardTransaction.MakeHash()
	block.Data = append(block.Data, rewardTransaction)
}

// todo: add a *Block.isValid() function
