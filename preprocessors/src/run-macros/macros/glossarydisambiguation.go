package macros

import (
	"log"
	"strings"
	"webdoky3/preprocessors/src/helpers"
)

func glossarydisambiguation(env *Environment, registry Registry, _ string) (string, error) {
	subItems := registry.GetSubItems(env.Path)
	content := "<dl>"
	for _, subItem := range subItems {
		log.Printf("subItem: %s", subItem.Chapter.Path)
		content += "<dt>[" + subItem.Chapter.Name + "](/" + strings.TrimSuffix(subItem.Chapter.Path, "/index.md") + ")</dt>"
		content += "<dd>" + helpers.GetSummary(&subItem, 0) + "</dd>"
	}
	content += "</dl>"
	return content, nil
}
