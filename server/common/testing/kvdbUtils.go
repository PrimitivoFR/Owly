package common_testing

import common_kvdb "primitivofr/owly/server-common/kvdb"

//SaveToKvdb allows to save to the testing kvdb in one line of code
func SaveToKvdb(bucketName string, key string, value string) {

	db, err := common_kvdb.InitDB("testingdb", bucketName)
	CheckErr(err, "Error from SaveToKvdb")
	err = common_kvdb.WriteData(db, bucketName, key, value)
	CheckErr(err, "Error from SaveToKvdb")
	db.Close()

}

// LoadFromKvdb allows to load from the testing kvdb in one line of code
func LoadFromKvdb(bucketName string, key string) string {

	db, err := common_kvdb.InitDB("testingdb", bucketName)
	CheckErr(err, "Error from LoadFromKvdb")

	accessToken, readallErr := common_kvdb.ReadData(db, bucketName, key)
	CheckErr(readallErr, "Error while reading for "+key)

	db.Close()

	return accessToken

}
