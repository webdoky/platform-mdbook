package helpers

import (
	"log"
	"os"
)

func DoesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}
