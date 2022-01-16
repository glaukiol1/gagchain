package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	Timestamp   int
	From        []byte
	PubKeyBytes []byte
	To          []byte
	Amount      int
	Signature   []byte
	Hash        [32]byte
}

func (bc *Blockchain) NewTransactionInstance(from *ecdsa.PublicKey, to string, amount int) *Transaction {
	return &Transaction{int(time.Now().Unix()), []byte(crypto.PubkeyToAddress(*from).Hex()), crypto.FromECDSAPub(from), []byte(to), amount, []byte{}, [32]byte{}}
}

func (trns *Transaction) GetTransactionJSON() string {
	dat, err := json.Marshal(Transaction{trns.Timestamp, trns.From, trns.PubKeyBytes, trns.To, trns.Amount, []byte{}, [32]byte{}})
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

func (trns *Transaction) IsValid() bool {
	hash := crypto.NewKeccakState()
	hash.Write(trns.PubKeyBytes[1:])
	println(string(trns.From), string(trns.To), fmt.Sprint(trns.Amount))
	if strings.ToLower(hexutil.Encode(hash.Sum(nil)[12:])) != strings.ToLower(string(trns.From)) {
		panic("Not matching PubKeyBytes with From")
	}
	return true
}

func (trns *Transaction) MakeHash() [32]byte {
	if len(trns.Signature) == 0 {
		panic("Set-up transaction before calling MakeHash")
	}
	dat, err := json.Marshal(trns)
	if err != nil {
		panic(err)
	}
	trns.Hash = sha256.Sum256(dat)
	return trns.Hash
}
