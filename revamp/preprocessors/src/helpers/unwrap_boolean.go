package preprocessor_helpers

func UnwrapBoolean(s string) bool {
	s = UnwrapString(s)
	if s == "false" || s == "0" || s == "null" {
		return false
	}
	return true
}
