package embedlivesample

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
)

func parseEmbedlivesampleArgs(args string) (string, string, string, string, error) {
	// Split the args string into a slice of strings
	// using the comma as the separator
	// (e.g., "termName,displayName" -> ["termName", "displayName"])
	argSlice := strings.Split(args, ",")
	var parentId, width, height, screenshotUrl string
	if len(argSlice) == 0 {
		return "", "", "", "", nil
	}
	switch len(argSlice) {
	case 6:
		fallthrough
	case 5:
		fallthrough
	case 4:
		screenshotUrl = helpers.UnwrapString(argSlice[3])
		fallthrough
	case 3:
		height = helpers.UnwrapString(argSlice[2])
		fallthrough
	case 2:
		width = helpers.UnwrapString(argSlice[1])
		fallthrough
	case 1:
		parentId = helpers.UnwrapString(argSlice[0])
	default:
		return "", "", "", "", errors.New("too many arguments")
	}
	parentId = helpers.GetSectionId(parentId)
	if height != "" {
		// If height is less than MIN_HEIGHT, set it to MIN_HEIGHT
		heightInt, err := strconv.Atoi(height)
		if err != nil {
			log.Println(err)
		} else {
			if heightInt < MIN_HEIGHT {
				height = strconv.Itoa(MIN_HEIGHT)
			}
		}
	}
	return parentId, width, height, screenshotUrl, nil
}
