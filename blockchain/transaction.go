package blockchain

import (
	"time"
)

type Transaction struct {
	Timestamp int
	From      []byte
	To        []byte
	Amount    int
}

func NewTransactionInstance(from, to string, amount int) *Transaction {
	return &Transaction{int(time.Now().Unix()), []byte(from), []byte(to), amount}
}
