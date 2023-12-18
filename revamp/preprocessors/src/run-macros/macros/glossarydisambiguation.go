package macros

import (
	"log"
	"strings"
	preprocessor_helpers "webdoky3/revamp/preprocessors/src/helpers"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func glossarydisambiguation(env *environment.Environment, r *registry.Registry, _ string) (string, error) {
	subItems := r.GetSubItems(env.Path)
	content := "<dl>"
	for _, subItem := range subItems {
		log.Printf("subItem: %s", subItem.Chapter.Path)
		content += "<dt>[" + subItem.Chapter.Name + "](/" + strings.TrimSuffix(subItem.Chapter.Path, "/index.md") + ")</dt>"
		content += "<dd>" + preprocessor_helpers.GetSummary(&subItem, 0) + "</dd>"
	}
	content += "</dl>"
	return content, nil
}
