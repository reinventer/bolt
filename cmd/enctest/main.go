package main

import (
	"github.com/reinventer/bolt"
	//"github.com/boltdb/bolt"
	//"time"
	"errors"
	"fmt"
)

func main() {
	db, err := bolt.Open(`db.bolt`, 0600, &bolt.Options{
		EncryptionKey: []byte(`keykeykeykeykey!`),
	})
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(`testbucket`))
		if err != nil {
			return err
		}
		return b.Put([]byte(`key`), []byte(`test`))
	})
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}

	db, err = bolt.Open(`db.bolt`, 0600, &bolt.Options{
		EncryptionKey: []byte(`keykeykeykeykey!`),
	})
	if err != nil {
		panic(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(`testbucket`))
		if b == nil {
			return errors.New(`bucket not found`)
		}
		fmt.Printf("result: %s\n", b.Get([]byte(`key`)))
		return nil
	})
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
