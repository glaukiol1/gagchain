package db

import (
	"fmt"

	"github.com/glaukiol1/gagchain/blockchain"
)

func (db *DB) UpdateDB() {
	tmp_bc := blockchain.Blockchain{db.ParseDB()} // ignore warning

	next_id := (tmp_bc.Blocks[len(tmp_bc.Blocks)-1].Id + 1)
	println("ID to get " + fmt.Sprint(next_id) + " until lastest block")
	//TODO
	// request all blocks since {next_id} from the selected node
	// and update the DB
}
