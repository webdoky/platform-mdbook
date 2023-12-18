package preprocessor_helpers

func DoesPageExist(slug string) bool {
	return DoesFileExist("./content/files/en-us/" + slug + "/index.md")
}
