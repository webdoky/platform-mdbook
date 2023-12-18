package preprocessor_helpers

import (
	"log"
	"os"
)

func GetFilesInFolder(folderPath string) []os.DirEntry {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	return files
}
