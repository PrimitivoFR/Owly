package common_kvdb

import (
	"os"

	"github.com/boltdb/bolt"
)

// InitDB init a boltdb object + db file + bucket if needed
// dbName is the name of the db (folder + .db file)
// bucketName is the name of the bucket that will contain data
func InitDB(dbName string, bucketName string) (*bolt.DB, error) {

	//ex, _ := os.Executable()

	//exPath := filepath.Dir(ex)
	//fmt.Println(exPath)

	if _, err := os.Stat("../../" + dbName); os.IsNotExist(err) {
		os.Mkdir("../../"+dbName, 0700)
	}

	db, err := bolt.Open("../../"+dbName+"/"+dbName+".db", 0644, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})

	return db, err

}
