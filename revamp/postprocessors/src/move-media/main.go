package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var SOURCE_FOLDER = "content/files/uk"
var TARGET_FOLDER = "book/uk/docs"

func copyFileToFolder(filePath string, folderPath string) error {
	filePathWithoutFolder := filePath[strings.LastIndex(filePath, "/"):]
	source, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(folderPath + filePathWithoutFolder)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func replaceInFile(filePath string, old string, new string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	fileAsString := string(file)
	fileAsString = strings.ReplaceAll(fileAsString, old, new)
	return os.WriteFile(filePath, []byte(fileAsString), 0644)
}

func processFile(path string, info os.FileInfo, err error) error {
	// log.Println("Processing " + path)
	if err != nil {
		return err
	}
	if info.Name() == "." || info.Name() == ".." {
		return nil
	}
	// fullFilePath := path + "/" + info.Name()
	if info.IsDir() {
		return nil
		// 	return filepath.Walk(fullFilePath, processFile)
	}
	if strings.HasSuffix(info.Name(), ".md") {
		return nil
	}
	fileFolder := path[:strings.LastIndex(path, "/")]
	frontmatterData, err := getFrontmatterData(fileFolder + "/index.md")
	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Println(err)
			return nil
		}
		return err
	}
	targetFolder := TARGET_FOLDER + "/" + frontmatterData.Slug
	err = os.MkdirAll(targetFolder, os.ModePerm)
	if err != nil {
		return err
	}
	log.Println("Moving " + path + " to " + targetFolder)
	err = copyFileToFolder(path, targetFolder)
	if err != nil {
		return err
	}
	return replaceInFile(targetFolder+"/index.html", "\""+info.Name()+"\"", "\"/uk/docs/"+frontmatterData.Slug+"/"+info.Name()+"\"")
}

func main() {
	err := filepath.Walk(SOURCE_FOLDER, processFile)
	if err != nil {
		log.Fatal(err)
	}
}
