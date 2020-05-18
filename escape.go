package jerr

import "strings"

// escapeJSON escapes the characters that are reserved in JSON.
func escapeJSON(s string) string {
	const newline = "\\n"
	var replacer = strings.NewReplacer(
		"\"", "\\\"",
		"\t", "\\t",
		"\r\n", newline,
		"\r", newline,
		"\n", newline,
	)
	return replacer.Replace(s)
}
