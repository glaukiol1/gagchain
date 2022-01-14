package blockchain

func (bc *Blockchain) GetBalance(address string, tp *TransactionPool) int {
	var balance int
	for _, block := range bc.Blocks {
		for _, trns := range block.Data {
			if string(trns.To) == address {
				balance += trns.Amount
			}
			if string(trns.From) == address {
				balance += -trns.Amount
			}
		}
	}
	for _, trns := range tp.pool {
		// TODO: make sure that the transaction is valid before adding to balance
		if string(trns.To) == address {
			balance += trns.Amount
		}
		if string(trns.From) == address {
			balance += -trns.Amount
		}
	}
	return balance
}
