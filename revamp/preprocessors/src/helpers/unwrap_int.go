package preprocessor_helpers

import (
	"log"
	"strconv"
)

func UnwrapInt(s string) int {
	s = UnwrapString(s)
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
