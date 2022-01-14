package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
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

func sign(privateKey *ecdsa.PrivateKey, data string) []byte {
	// privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	// if err != nil {
	// log.Fatal(err)
	// }

	hash := crypto.Keccak256Hash([]byte(data))

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return signature
}

func VerifySign(publicKeyECDSA *ecdsa.PublicKey, signature []byte, dat string) bool {

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte(dat)
	hash := crypto.Keccak256Hash(data)

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	return matches
}

func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {
	privateKeyBytes := crypto.FromECDSA(privateKey)
	return hexutil.Encode(privateKeyBytes)[2:]
}
