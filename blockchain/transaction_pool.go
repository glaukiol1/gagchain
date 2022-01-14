package blockchain

// ability to;
// --- add transactions
// --- mine transactions
// --- --- Part of this is checking the signature
// --- sort transactions (by gas fee)

const MIN_transactions_in_block = 2 // this is excluding mint address

type TransactionPool struct {
	pool []*Transaction
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{}
}
