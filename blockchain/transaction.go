package blockchain

import (
	"time"
)

type Transaction struct {
	Timestamp int
	Input     []byte
	Output    []byte
}

func NewTransactionInstance(input, output []byte) *Transaction {
	return &Transaction{int(time.Now().Unix()), input, output}
}
