package preprocessor_helpers

import (
	"log"
	"os"
)

func WriteToStderr(message string) {
	_, err := os.Stderr.WriteString(message)
	if err != nil {
		log.Fatal(err)
	}
}
