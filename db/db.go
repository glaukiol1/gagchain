package db

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/glaukiol1/gagchain/blockchain"
)

type DB struct {
	location string
}

func (db *DB) GetContents() string {
	if db.location == "" {
		panic("Set up the DB first")
	}
	dat, err := os.ReadFile(db.location)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

func (db *DB) Write(data string) {
	dat := []byte(data)
	err := os.WriteFile(db.location, dat, 0644)
	if err != nil {
		panic(err)
	}
}

func (db *DB) ParseDB() []*blockchain.Block {
	var v []*blockchain.Block
	json.Unmarshal([]byte(db.GetContents()), &v)
	return v
}

func GetDB(location string) DB {
	return DB{location}
}

func DB_DoesExist(location string) bool {
	if _, err := os.Stat(location); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
