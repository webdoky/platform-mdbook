package cssinfo

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func initial(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	initialValue := data.Initial
	if len(initialValue) == 1 {
		singleInitial := initialValue[0]
		result, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
			Text: singleInitial,
		})
		if err != nil {
			return "", err
		}
		return result, nil
	} else {
		return as_longhands(env, reg, initialValue)
	}
}
