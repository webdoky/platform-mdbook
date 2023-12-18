package macros

import (
	"errors"
	"html/template"
	"log"
	"strings"
	preprocessor_helpers "webdoky3/revamp/preprocessors/src/helpers"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/preprocessor"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"

	"golang.org/x/exp/slices"
)

func parseListsubpagesArgs(args string) (string, int, bool, bool, error) {
	argSlice := strings.Split(args, ",")
	if len(argSlice) == 0 {
		return "", 0, false, false, errors.New("no arguments")
	}
	var path string
	if len(argSlice) >= 1 {
		path = preprocessor_helpers.UnwrapString(argSlice[0])
	}
	depth := 1
	if len(argSlice) >= 2 {
		depth = preprocessor_helpers.UnwrapInt(argSlice[1])
	}
	var reverse, ordered bool
	if len(argSlice) >= 3 {
		reverse = preprocessor_helpers.UnwrapBoolean(argSlice[2])
	}
	if len(argSlice) >= 4 {
		ordered = preprocessor_helpers.UnwrapBoolean(argSlice[3])
	}
	log.Printf("path: %q, depth: %d, reverse: %v, ordered: %v", path, depth, reverse, ordered)
	return path, depth, reverse, ordered, nil
}

func listsubpages(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	path, depth, reverse, ordered, err := parseListsubpagesArgs(args)
	if err != nil {
		return "", err
	}
	if path == "" {
		path = env.Path
	}
	path = strings.TrimPrefix(path, "/")
	if strings.HasPrefix(path, "en-US/") {
		path = env.Locale + path[len("en-US"):]
	}
	// log.Println("path:", path)
	subItems := reg.GetSubItems(path)
	// log.Printf("Found %d subitems", len(subItems))
	if len(subItems) == 0 {
		return "", nil
	}
	if reverse {
		slices.Reverse(subItems)
	}
	return renderSubpages(subItems, depth, ordered)
}

func renderSubpages(subItems []preprocessor.Section, depth int, ordered bool) (string, error) {
	if depth == 0 {
		// log.Println("depth == 0")
		return "", nil
	}
	itemsHtml := make([]string, len(subItems))
	for i, item := range subItems {
		link, err := renderhtml.RenderA(&renderhtml.AParams{
			Href: "/" + strings.TrimSuffix(item.Chapter.Path, "/index.md"),
			Text: item.Chapter.Name,
		})
		// log.Println(link)
		if err != nil {
			return "", err
		}
		subListHtml, err := renderSubpages(subItems, depth-1, ordered)
		// log.Println(subListHtml)
		if err != nil {
			return "", err
		}
		html, err := renderhtml.RenderLi(&renderhtml.LiParams{
			InnerHtml: template.HTML(link + subListHtml),
		})
		// log.Println(html)
		if err != nil {
			return "", err
		}
		itemsHtml[i] = html
	}
	var outerHtml string
	var err error
	if ordered {
		params := renderhtml.OlParams{
			InnerHtml: template.HTML(strings.Join(itemsHtml, "")),
		}
		outerHtml, err = renderhtml.RenderOl(&params)
	} else {
		params := renderhtml.UlParams{
			InnerHtml: template.HTML(strings.Join(itemsHtml, "")),
		}
		outerHtml, err = renderhtml.RenderUl(&params)
	}
	if err != nil {
		return "", err
	}
	return outerHtml, nil
}
