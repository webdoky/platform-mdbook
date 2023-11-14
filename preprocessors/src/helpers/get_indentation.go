package helpers

func GetIndentation(depth int) string {
	indentation := ""
	for i := 0; i < depth; i++ {
		indentation += "  "
	}
	return indentation
}
