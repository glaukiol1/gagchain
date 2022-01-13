package db

import "os"

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

func GetDB(location string) DB {
	return DB{location}
}
