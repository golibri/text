package text

/*
 * checks: methods that inspect the string and return a bool
 */

import "strings"
import "regexp"

// DoesContainString checks whether a given string is contained in the text.
func (t Text) DoesContainString(s string) bool {
	return strings.Contains(t.ToString(), s)
}

// DoesMatchPattern tests a gioven regex string against the text
func (t Text) DoesMatchPattern(s string) bool {
	return regexp.MustCompile(s).Match(t.ToBytes())
}

// EndsWith is similar to StartsWith, but checks the last characters
func (t Text) EndsWith(pattern string) bool {
	return strings.HasSuffix(t.ToString(), pattern)
}

// IsEqual compares a text directly with a given string
func (t Text) IsEqual(s string) bool {
	return t == New(s)
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
