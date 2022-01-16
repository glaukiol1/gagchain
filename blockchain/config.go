package blockchain

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type Mining_Minting_Address struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

type myAddress struct {
	privateKey    *ecdsa.PrivateKey
	publicKey     *ecdsa.PublicKey
	publicAddress string
}

var privateHex = "99fb794cc33697f7f1d3de7ad55b8eca2fde97c7f7fb5d1b0c0de1dd757c8371"

var _privateKey, _ = crypto.HexToECDSA(privateHex)

var _publicKey = &_privateKey.PublicKey

var MintAddress = Mining_Minting_Address{_privateKey, _publicKey}

var pk, _ = crypto.HexToECDSA("d2d9a9aa0fce4a8a1d3141300fa8a0c0087f7ae93dd396d5198381c440584361") // 0x4390B0820B4257d8936759e5e043e91a1F9E0BeC
var pb = &pk.PublicKey
var MyAddress = myAddress{pk, pb, PubkeyToAddress(pb)}

const Mining_Node = true

const Reward = 100

var Difficulty = 15 // miners on the network / more power
