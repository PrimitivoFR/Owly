package common_kvdb

import (
	"fmt"

	"github.com/boltdb/bolt"
)

// ReadData reads any value from boltdb, using a key
// bucketName is the name of the bucket that contains data
func ReadData(db *bolt.DB, bucketName string, key string) (string, error) {

	var res []byte

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		res = bucket.Get([]byte(key))

		return nil
	})

	if err != nil {
		return "", err
	}

	return string(res), nil
}

// ReadAll reads all data from any boltdb
// bucketName is the name of the bucket that contains data
func ReadAll(db *bolt.DB, bucketName string) (map[string]string, error) {

	data := make(map[string]string)

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			val := fmt.Sprintf("%s", v)
			key := fmt.Sprintf("%s", k)
			data[key] = val
		}
		return nil
	})

	return data, err
}
