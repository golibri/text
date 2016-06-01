// Package text wraps the default string type and adds useful methods, OOP style
package text

import "strings"
import "regexp"

// Text is the same as a string, comparison operators etc apply as usual
type Text string

/*
 * unclassified methods, useful stuff
 */

// SplitString splits the string on every occurence of a given pattern
func (t Text) SplitString(pattern string) []string {
	return strings.Split(t.ToString(), pattern)
}

// SplitPattern is the same as SplitString, but uses regex pattern as a string
func (t Text) SplitPattern(pattern string) []string {
	return regexp.MustCompile(pattern).Split(t.ToString(), -1)
}
