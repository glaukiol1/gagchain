package blockchain

func ValidateBlock(block *Block, bc *Blockchain) *Block {
	if block.Id == 0 { // if the block is the genesis block
		return block
	} else {
		oldTransactions := block.Data
		var newTransactions []*Transaction // these are verified
		newBlock := &Block{[]byte{}, newTransactions, block.PrevHash, -1, block.Id, block.Timestamp}
		for _, oldTrns := range oldTransactions {
			if oldTrns.IsValid(&TransactionPool{}) && oldTrns.IsReady() && !(bc.GetBalance(string(oldTrns.From), &TransactionPool{newTransactions /* this is safe, since {newTransaction} is only verified transactions */}) < oldTrns.Amount) && oldTrns.VerifySignature() {
				newTransactions = append(newTransactions, oldTrns)
			} else {
				println("Invalid transaction")
			}
		}
		newBlock.Data = newTransactions
		return newBlock
	}

}
