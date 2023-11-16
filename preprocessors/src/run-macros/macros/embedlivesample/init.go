package embedlivesample

import (
	"html/template"
	"log"
	"os"
	"strings"
)

func init() {
	var err error
	tLiveSample, err = template.ParseFiles("./preprocessors/src/run-macros/macros/embedlivesample/templates/livesample.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tLiveSample, err = template.ParseFiles("./templates/livesample.tmpl")
	}
	if err != nil {
		log.Println(os.Getwd())
		log.Fatal(err)
	}
	tEmbedLiveSample, err = template.ParseFiles("./preprocessors/src/run-macros/macros/embedlivesample/templates/embedlivesample.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tEmbedLiveSample, err = template.ParseFiles("./templates/embedlivesample.tmpl")
	}
	if err != nil {
		log.Println(os.Getwd())
		log.Fatal(err)
	}
}
