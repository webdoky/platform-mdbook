package environment

import "webdoky3/revamp/helpers"

type Environment struct {
	Content     *string
	Frontmatter *helpers.FrontmatterData
	Locale      string
	Path        string
	Summary     string
}
