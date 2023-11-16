package embedlivesample

import (
	"html/template"
	"log"
	"strings"
	"webdoky3/preprocessors/src/run-macros/environment"
	"webdoky3/preprocessors/src/run-macros/registry"
)

var tLiveSample *template.Template

type CodeItem struct {
	Code     string
	Language string
}

type LiveSampleParams struct {
	Css   []template.CSS
	Html  []template.HTML
	Js    []template.JS
	Title string
}

var MIN_HEIGHT = 60

func Embedlivesample(env *environment.Environment, _ *registry.Registry, args string) (string, error) {
	log.Printf("embedlivesample: %s", args)
	log.Println("env.Path:", env.Path)
	parentId, width, height, _, err := parseEmbedlivesampleArgs(args)
	if err != nil {
		return "", err
	}
	targetSection, sectionTitle, err := extractTargetSection(*env.Content, parentId)
	if err != nil {
		return "", err
	}
	codeItems := extractCodeItems(targetSection)
	var cssItems []template.CSS
	var htmlItems []template.HTML
	var jsItems []template.JS
	for _, codeItem := range codeItems {
		switch codeItem.Language {
		case "css":
			cssItems = append(cssItems, template.CSS(codeItem.Code))
		case "html":
			htmlItems = append(htmlItems, template.HTML(codeItem.Code))
		case "js":
			jsItems = append(jsItems, template.JS(codeItem.Code))
		}
	}
	liveSampleParams := LiveSampleParams{
		Css:   cssItems,
		Html:  htmlItems,
		Js:    jsItems,
		Title: sectionTitle,
	}
	exampleFolder := "./live-samples/" + strings.TrimSuffix(env.Path, "/index.md")
	examplePath := exampleFolder + "/" + parentId + ".html"
	err = saveLiveSample(exampleFolder, examplePath, &liveSampleParams)
	if err != nil {
		return "", err
	}
	embedLiveSampleParams := EmbedLiveSampleParams{
		Class:  "sample-code-frame",
		Height: template.HTMLAttr(height),
		Id:     template.HTMLAttr(parentId),
		Title:  sectionTitle,
		Url:    template.URL(strings.TrimPrefix(examplePath, ".")),
		Width:  template.HTMLAttr(width),
	}
	return renderEmbedlivesample(&embedLiveSampleParams)
}
