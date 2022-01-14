package blockchain

// ability to;
// --- add transactions
// --- mine transactions
// --- edit transaction
// --- --- Part of this is checking the signature
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
	tp.pool[pos] = transaction
	return pos
}

func (tp *TransactionPool) EditTransactoin(transactionPos int, newTransaction *Transaction) {
	tp.pool[transactionPos] = newTransaction
}
