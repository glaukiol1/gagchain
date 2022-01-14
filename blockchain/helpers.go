package blockchain

import "strings"

func (bc *Blockchain) GetBalance(address string, tp *TransactionPool) int {
	var balance int
	address = strings.ToLower(address)
	for _, block := range bc.Blocks {
		for _, trns := range block.Data {
			if strings.ToLower(string(trns.To)) == address {
				balance += trns.Amount
			}
			if strings.ToLower(string(trns.From)) == address {
				balance += -trns.Amount
			}
		}
	}
	for _, trns := range tp.pool {
		if strings.ToLower(string(trns.To)) == address {
			balance += trns.Amount
		}
		if strings.ToLower(string(trns.From)) == address {
			balance += -trns.Amount
		}
	}
	return balance
}
