package preprocessor

import (
	"log"
	"os"
	"strings"
	"time"
)

func LogToFile() *os.File {
	if os.Getenv("PRODUCTION") == "true" {
		log.SetOutput(os.Stderr)
		return nil
	}
	appName := os.Args[0]
	slashIndex := strings.LastIndex(appName, "/")
	if slashIndex != -1 {
		appName = appName[slashIndex+1:]
	}
	err := os.MkdirAll("logs/"+appName, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating logs directory: %v", err)
	}
	f, err := os.OpenFile("logs/"+appName+"/"+time.Now().Format("2006-01-02T15:04:05")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	return f
}
