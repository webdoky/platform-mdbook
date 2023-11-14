package helpers

import "os"

func WriteToStderr(message string) {
	_, err := os.Stderr.WriteString(message)
	if err != nil {
		panic(err)
	}
}
