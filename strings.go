package utils

import "strings"

// MakeStrLookupMap create a lookup map for a string slice.
func MakeStrLookupMap(xs []string) map[string]struct{} {
	m := make(map[string]struct{}, 0)
	for _, x := range xs {
		m[x] = struct{}{}
	}
	return m
}

// CollapseWhitespace removes spaces around and between words and
// returns just the words, with a single space between
func CollapseWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}
