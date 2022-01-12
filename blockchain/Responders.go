package blockchain

import (
	"encoding/json"
)

// responders to different input messages

func InitResponders(bc Blockchain) {

}

func (bc *Blockchain) RequestChain() string {
	out, err := json.Marshal(bc.Blocks)
	if err != nil {
		panic(err)
	}
	return string(out)
}
