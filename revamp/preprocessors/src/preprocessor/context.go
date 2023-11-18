package preprocessor

type ContextBookConfig struct {
	Authors      []string `json:"authors"`
	Language     string   `json:"language"`
	Multilingual bool     `json:"multilingual"`
	Source       string   `json:"src"`
	Title        string   `json:"title"`
}
type ContextBuildConfig struct {
	BuildDir                string   `json:"build-dir"`
	CreateMissing           bool     `json:"create-missing"`
	ExtraWatchDirs          []string `json:"extra-watch-dirs"`
	UseDefaultPreprocessors bool     `json:"use-default-preprocessors"`
}
type ContextOutputHtmlConfig struct {
	AdditionalCss []string `json:"additional-css"`
	Fold          struct {
		Enable bool `json:"enable"`
	} `json:"fold"`
	GitRepositoryUrl string `json:"git-repository-url"`
	NoSectionLabel   bool   `json:"no-section-label"`
	Print            struct {
		Enable bool `json:"enable"`
	} `json:"print"`
	Search struct {
		Enable        bool `json:"enable"`
		UseBooleanAnd bool `json:"use-boolean-and"`
	} `json:"search"`
}
type ContextOutputConfig struct {
	Html ContextOutputHtmlConfig `json:"html"`
}
type ContextConfig struct {
	Book   ContextBookConfig   `json:"book"`
	Build  ContextBuildConfig  `json:"build"`
	Output ContextOutputConfig `json:"output"`
}
type Context struct {
	Config        ContextConfig `json:"config"`
	MdbookVersion string        `json:"mdbook_version"`
	Renderer      string        `json:"renderer"`
	Root          string        `json:"root"`
}

func ConvertMapToContext(data map[string]interface{}) *Context {
	var context Context
	context.Config.Book.Source = data["config"].(map[string]interface{})["book"].(map[string]interface{})["src"].(string)
	context.Root = data["root"].(string)
	return &context
}
