package dbclient

import (
	"github.com/boltdb/bolt"
	"log"
)

type IBoltClient interface {
	OpenBoltDb()
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("carts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}
