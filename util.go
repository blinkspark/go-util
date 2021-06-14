package util

import (
	"log"
	"os"
)

// CheckErr check the error
func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// just ignore everything
func Ignore(any ...interface{}) {
	log.Printf("ignoring %#+v\n", any)
}

// PathExists is file exist
func PathExists(path string) bool {
	_, err := os.Lstat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
