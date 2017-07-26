package db

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

//DBPATH is path of db
const DBPATH = "./db/bolt.db"

//BUCKET is db name equivalent
var BUCKET = []byte("routes")

func GetData(key []byte) (string, error) {
	db, err := bolt.Open(DBPATH, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var value string
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BUCKET)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", BUCKET)
		}

		val := bucket.Get(key)
		value = string(val)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return value, err
}

func SetData(key []byte, value []byte) error {
	db, err := bolt.Open(DBPATH, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(BUCKET)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
