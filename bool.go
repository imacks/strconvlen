package strconvlen

// Bool returns the same result as len(strconv.FormatBool(b)).
func Bool(b bool) int {
	if b {
		return 4 // true
	}
	return 5 // false
}
