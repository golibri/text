package text

/*
 * transformations: chainable methods that change the string and return a new one
 */

import "strings"
import "regexp"

// Capitalize uppercases the first letter of every word
func (t Text) Capitalize() Text {
	return New(strings.Title(string(t)))
}

// Downcase lowercases all characters in the string
func (t Text) Downcase() Text {
	return New(strings.ToLower(t.ToString()))
}

// ReplaceString replaces all occurences of `old` with `new`
func (t Text) ReplaceString(old string, new string) Text {
	return Text(strings.Replace(t.ToString(), old, new, -1))
}

// ReplacePattern accepts a regex pattern as a string, same as ReplaceString()
func (t Text) ReplacePattern(pattern string, new string) Text {
	return Text(regexp.MustCompile(pattern).ReplaceAllString(t.ToString(), new))
}

// Reverse does exactly this, char by char
func (t Text) Reverse() Text {
	chars := t.ToChars()
	length := len(chars)
	result := make([]string, length)
	for i := length - 1; i >= 0; i-- {
		result[length-i-1] = chars[i]
	}
	return Text(strings.Join(result, ""))
}

// Strip removes all whitespace at the beginning and end of the string
func (t Text) Strip() Text {
	return t.StripLeft().StripRight()
}

// StripLeft removes all whitespace at the beginning of the string
func (t Text) StripLeft() Text {
	return t.ReplacePattern(`^\s+`, "")
}

// StripRight removes all whitespace at the end of the string
func (t Text) StripRight() Text {
	return t.ReplacePattern(`\s+$`, "")
}

// Upcase replaces every character in the string with its uppercase variant
func (t Text) Upcase() Text {
	return New(strings.ToUpper(t.ToString()))
}
