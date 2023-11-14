package helpers

import (
	"log"
	"os"
	"strings"

	"github.com/adrg/frontmatter"
)

type formatterData1 struct {
	BrowserCompat string   `yaml:"browser-compat"`
	PageType      string   `yaml:"page-type"`
	Slug          string   `yaml:"slug"`
	SpecUrls      string   `yaml:"spec-urls"`
	Status        []string `yaml:"status"`
	Title         string   `yaml:"title"`
}

type formatterData2 struct {
	BrowserCompat []string `yaml:"browser-compat"`
	PageType      string   `yaml:"page-type"`
	Slug          string   `yaml:"slug"`
	SpecUrls      string   `yaml:"spec-urls"`
	Status        []string `yaml:"status"`
	Title         string   `yaml:"title"`
}
type formatterData3 struct {
	BrowserCompat string   `yaml:"browser-compat"`
	PageType      string   `yaml:"page-type"`
	Slug          string   `yaml:"slug"`
	SpecUrls      []string `yaml:"spec-urls"`
	Status        []string `yaml:"status"`
	Title         string   `yaml:"title"`
}

type FrontmatterData struct {
	BrowserCompat []string `yaml:"browser-compat"`
	PageType      string   `yaml:"page-type"`
	Slug          string   `yaml:"slug"`
	SpecUrls      []string `yaml:"spec-urls"`
	Status        []string `yaml:"status"`
	Title         string   `yaml:"title"`
}

func adjustFormat1(frontmatterData *formatterData1) *FrontmatterData {
	frontmatterData1 := &FrontmatterData{
		BrowserCompat: []string{frontmatterData.BrowserCompat},
		PageType:      frontmatterData.PageType,
		Slug:          frontmatterData.Slug,
		SpecUrls:      []string{frontmatterData.SpecUrls},
		Status:        frontmatterData.Status,
		Title:         frontmatterData.Title,
	}
	return frontmatterData1
}

func adjustFormat2(frontmatterData *formatterData2) *FrontmatterData {
	frontmatterData2 := &FrontmatterData{
		BrowserCompat: frontmatterData.BrowserCompat,
		PageType:      frontmatterData.PageType,
		Slug:          frontmatterData.Slug,
		SpecUrls:      []string{frontmatterData.SpecUrls},
		Status:        frontmatterData.Status,
		Title:         frontmatterData.Title,
	}
	return frontmatterData2
}

func adjustFormat3(frontmatterData *formatterData3) *FrontmatterData {
	frontmatterData3 := &FrontmatterData{
		BrowserCompat: []string{frontmatterData.BrowserCompat},
		PageType:      frontmatterData.PageType,
		Slug:          frontmatterData.Slug,
		SpecUrls:      frontmatterData.SpecUrls,
		Status:        frontmatterData.Status,
		Title:         frontmatterData.Title,
	}
	return frontmatterData3
}

func GetFrontmatterData(markdownFilePath string) (*FrontmatterData, error) {
	markdown, err := os.ReadFile(markdownFilePath)
	if err != nil {
		return nil, err
	}

	var frontmatterData FrontmatterData
	var frontmatterData1Value formatterData1
	_, err = frontmatter.Parse(strings.NewReader(string(markdown)), &frontmatterData1Value)
	if err == nil {
		log.Printf("%s has format 1\n", markdownFilePath)
		return adjustFormat1(&frontmatterData1Value), nil
	}
	// if err message includes "cannot unmarshal !!str"
	if strings.Contains(err.Error(), "cannot unmarshal !!seq") {
		var frontmatterData2Value formatterData2
		_, err = frontmatter.Parse(strings.NewReader(string(markdown)), &frontmatterData2Value)
		if err == nil {
			log.Printf("%s has format 2\n", markdownFilePath)
			return adjustFormat2(&frontmatterData2Value), nil
		}
	}
	if strings.Contains(err.Error(), "cannot unmarshal !!str") || strings.Contains(err.Error(), "cannot unmarshal !!seq") {
		var frontmatterData3Value formatterData3
		_, err = frontmatter.Parse(strings.NewReader(string(markdown)), &frontmatterData3Value)
		if err == nil {
			log.Printf("%s has format 3\n", markdownFilePath)
			return adjustFormat3(&frontmatterData3Value), nil
		}
	}
	if strings.Contains(err.Error(), "cannot unmarshal !!str") || strings.Contains(err.Error(), "cannot unmarshal !!seq") {
		_, err = frontmatter.Parse(strings.NewReader(string(markdown)), &frontmatterData)
		if err == nil {
			log.Printf("%s has format 4\n", markdownFilePath)
			return &frontmatterData, nil
		}
	}
	log.Printf("Error parsing frontmatter for %s: %s", markdownFilePath, err)

	return nil, err
}
