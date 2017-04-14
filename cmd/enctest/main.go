package main

import (
	"github.com/reinventer/bolt"
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
}
