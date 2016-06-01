package text

import "strings"

/*
 * Casts / Export Utilities to various types
 */

// ToBytes returns the string as a byte slice
func (t Text) ToBytes() []byte {
	return []byte(t)
}

// ToChars returns a slice of characters of the underlying string
func (t Text) ToChars() []string {
	return strings.Split(t.ToString(), "")
}

// ToLines returns a string slice with all the different lines
func (t Text) ToLines() []string {
	return t.
		ReplaceString("\r", "\n").
		ReplaceString("\n\n", "\n").
		SplitString("\n")
}

// ToString returns a plain string type
func (t Text) ToString() string {
	return string(t)
}
