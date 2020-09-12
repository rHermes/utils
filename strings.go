package utils

import "strings"

// CollapseWhitespace removes spaces around and between words and
// returns just the words, with a single space between
func CollapseWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}
