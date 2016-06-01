package text

/*
 * checks: methods that inspect the string and return a bool
 */

import "strings"

// EndsWith is similar to StartsWith, but checks the last characters
func (t Text) EndsWith(pattern string) bool {
	return strings.HasSuffix(t.ToString(), pattern)
}

// IsASCIIOnly checks if the string contain unicode characters
func (t Text) IsASCIIOnly() bool {
	return len(t.ToString()) == len(t.ToChars())
}

// IsEmpty checks whether the given string contains zero characters
func (t Text) IsEmpty() bool {
	return t.ToString() == ""
}

// StartsWith checks whether the strings first characters match the given pattern
func (t Text) StartsWith(pattern string) bool {
	return strings.HasPrefix(t.ToString(), pattern)
}
