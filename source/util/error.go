package util

import (
	"log"
)

// CheckError check and output errors
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
