package common_kvdb

import (
	"github.com/boltdb/bolt"
)

// WriteData writes data to a bolt db, with key and value
// bucketName is the name of the bucket that contains data
func WriteData(db *bolt.DB, bucketName string, key string, value string) error {

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		bucket.Put([]byte(key), []byte(value))

		return nil
	})

	return err

}
