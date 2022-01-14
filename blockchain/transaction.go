package blockchain

import (
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	Timestamp int
	From      []byte
	To        []byte
	Amount    int
	Signature []byte
}

func NewTransactionInstance(from, to string, amount int) *Transaction {
	return &Transaction{int(time.Now().Unix()), []byte(from), []byte(to), amount, []byte{}}
}

func (trns *Transaction) GetTransactionJSON() string {
	dat, err := json.Marshal(Transaction{trns.Timestamp, trns.From, trns.To, trns.Amount, []byte{}})
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func (trns *Transaction) Sign(privateHex string) {
	privateKey, err := crypto.HexToECDSA(privateHex)
	if err != nil {
		panic(err)
	}

	trns.Signature = sign(privateKey, trns.GetTransactionJSON())
}
