package common_testing

import "log"

// CheckErr checks error in a single line of code.
func CheckErr(e error, desc string) {
	if e != nil {
		log.Println(desc)
		panic(e)
	}
}
