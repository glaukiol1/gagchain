package blockchain

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

// sign messages
// verify signature

func Keygen() (*ecdsa.PublicKey, *ecdsa.PrivateKey) {
	privateKey, err := crypto.GenerateKey()
	publicKey := privateKey.Public()
	if err != nil {
		log.Fatal(err)
	}
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	return publicKeyECDSA, privateKey
}

func PubkeyToAddress(publicKeyECDSA *ecdsa.PublicKey) string {
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}
