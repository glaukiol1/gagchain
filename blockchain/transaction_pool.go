package blockchain

import "fmt"

// ability to;
// --- add transactions
// --- mine transactions
// --- --- Part of this is checking the signature
// --- edit transaction
// --- sort transactions (by gas fee)

const MIN_transactions_in_block = 2 // this is excluding mint address

type TransactionPool struct {
	pool []*Transaction
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{}
}

func (tp *TransactionPool) AddTransaction(transaction *Transaction) int {
	pos := len(tp.pool) // same as len(tp.pool)-1+1
	tp.pool = append(tp.pool, transaction)
	return pos
}

func (tp *TransactionPool) EditTransaction(transactionPos int, newTransaction *Transaction) {
	tp.pool[transactionPos] = newTransaction
}

func (tp *TransactionPool) MineTransactions() []*Transaction {
	var blockData []*Transaction
	if len(tp.pool) < MIN_transactions_in_block {
		return []*Transaction{}
	}
	for i := 0; i < len(tp.pool); i++ {
		trns := tp.pool[i]
		if !(trns.IsReady()) {
			// do something if it isnt ready
		} else {
			if !(trns.VerifySignature()) {
				// do something if the signature is invalid
				fmt.Printf("Transaction %s has an invalid signature", trns.GetTransactionJSON())
			} else {
				// mine transaction
				trns.MakeHash()
				blockData = append(blockData, trns)
			}
		}
	}
	return blockData
}
