package embedlivesample

import (
	"bytes"
	"html/template"
)

// initialized in init.go
var tEmbedLiveSample *template.Template

type EmbedLiveSampleParams struct {
	Class  template.HTMLAttr
	Height template.HTMLAttr
	Id     template.HTMLAttr
	Title  string
	Url    template.URL
	Width  template.HTMLAttr
}

func renderEmbedlivesample(params *EmbedLiveSampleParams) (string, error) {
	var embedLiveSampleBuffer bytes.Buffer
	err := tEmbedLiveSample.Execute(&embedLiveSampleBuffer, params)
	if err != nil {
		return "", err
	}
	return embedLiveSampleBuffer.String(), nil
}
