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
		if trns.IsReady() {
			if string(trns.To) == address {
				balance += trns.Amount
			}
			if string(trns.From) == address {
				balance += -trns.Amount
			}
		}

	}
	return balance
}
