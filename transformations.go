package text

/*
 * transformations: chainable text modifications (immutable)
 */

import "strings"
import "regexp"

// Capitalize uppercases the first letter of every word
func (t Text) Capitalize() Text {
	return New(strings.Title(string(t)))
}

// DeleteString Removes every occurence of the given string from the text
func (t Text) DeleteString(s string) Text {
	return t.ReplaceString(s, "")
}

// DeletePattern Removes every occurence of the given regex pattern from the text
func (t Text) DeletePattern(s string) Text {
	return t.ReplacePattern(s, "")
}

// Downcase lowercases all characters in the string
func (t Text) Downcase() Text {
	return New(strings.ToLower(t.ToString()))
}

// First truncates the text to the first `i` characters
func (t Text) First(i int) Text {
	if i < 1 {
		return New("")
	}
	if i >= t.Length() {
		return t
	}
	result := []string{}
	t.EachChar(func(c string) {
		if len(result) < i {
			result = append(result, c)
		}
	})
	return New(strings.Join(result, ""))
}

// Last truncates the text to the last `i` characters
func (t Text) Last(i int) Text {
	if i < 1 {
		return New("")
	}
	if i >= t.Length() {
		return t
	}
	return t.Reverse().First(i).Reverse()
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

// Upcase replaces every character in the string with its uppercase variant
func (t Text) Upcase() Text {
	return New(strings.ToUpper(t.ToString()))
}
