package preprocessor_helpers

import "strings"

// Remove YAML frontmatter from Markdown code
func RemoveFrontmatter(markdown string) string {
	frontmatterStart := strings.Index(markdown, "---")
	if frontmatterStart == -1 {
		return markdown
	}
	frontmatterEnd := strings.Index(markdown[frontmatterStart+3:], "---")
	if frontmatterEnd == -1 {
		return markdown
	}
	frontmatterEnd += frontmatterStart + 3
	return markdown[:frontmatterStart] + markdown[frontmatterEnd:]
}
