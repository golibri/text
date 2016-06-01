package text

/*
 * metrics: methods that inspect the string and return useful data about it
 */

import "strings"

// ByteSize returns the length of bytes of the given string
func (t Text) ByteSize() int {
	return len(t.ToBytes())
}

// Count determines how often a given phrase is contained in the string
func (t Text) Count(phrase string) int {
	return strings.Count(t.ToString(), phrase)
}

// Length returns the number of characters of the string
func (t Text) Length() int {
	return len(t.ToChars())
}
