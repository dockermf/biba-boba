package dbapi

// TODO: add query templates

// Processes user input. Forbidden characters are prepended by backslash.
func SanitizeString(input string) string {
	/* 1. Check if input is a valid string (prepend \ to special characters to
	 * avoid SQLi, length to MAX of the smallest fields db can store)
	 */
	return ""
}
