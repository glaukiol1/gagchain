package blockchain

import (
	"crypto/ecdsa"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	Timestamp   int
	From        []byte
	PubKeyBytes []byte
	To          []byte
	Amount      int
	Signature   []byte
}

func NewTransactionInstance(from *ecdsa.PublicKey, to string, amount int) *Transaction {
	return &Transaction{int(time.Now().Unix()), []byte(crypto.PubkeyToAddress(*from).Hex()), crypto.FromECDSAPub(from), []byte(to), amount, []byte{}}
}

func (trns *Transaction) GetTransactionJSON() string {
	dat, err := json.Marshal(Transaction{trns.Timestamp, trns.From, trns.PubKeyBytes, trns.To, trns.Amount, []byte{}})
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

func (trns *Transaction) VerifySignature() bool {
	publicKey, err := crypto.UnmarshalPubkey(trns.PubKeyBytes)
	if err != nil {
		panic(err)
	}

	return VerifySign(publicKey, trns.Signature, trns.GetTransactionJSON())
}

func (trns *Transaction) IsReady() bool {
	if len(trns.Signature) == 0 {
		return false
	}
	return true
}
