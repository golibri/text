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
	if i >= t.Length() {
		return t
	}
	return t.Reverse().First(i).Reverse()
}

// LeftJust increases the length of the string to `i` characters, preps whitespace to left
func (t Text) LeftJust(length int) Text {
	return t.LeftJustString(length, " ")
}

// LeftJustString increases the length of the string to `i` characters, preps `seq` to left
func (t Text) LeftJustString(length int, seq string) Text {
	tLen := t.Length()
	if tLen >= length {
		return t
	}
	seqLen := New(seq).Length()
	str := strings.Repeat(seq, (length-tLen)/seqLen) +
		New(seq).First((length-tLen)%seqLen).ToString() +
		t.ToString()
	return New(str)
}

// LeftPad prepends `i` whitespaces to the string
func (t Text) LeftPad(i int) Text {
	return t.LeftPadString(i, " ")
}

// LeftPadString prepends `i` times the given sequence to a string
func (t Text) LeftPadString(i int, seq string) Text {
	return New(strings.Repeat(seq, i) + t.ToString())
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

// RightJust increases the length of the string to `i` characters, appends whitespace to left
func (t Text) RightJust(length int) Text {
	return t.RightJustString(length, " ")
}

// RightJustString increases the length of the string to `i` characters, app `seq` to left
func (t Text) RightJustString(length int, seq string) Text {
	tLen := t.Length()
	if tLen >= length {
		return t
	}
	seqLen := New(seq).Length()
	str := t.ToString() +
		strings.Repeat(seq, (length-tLen)/seqLen) +
		New(seq).First((length-tLen)%seqLen).ToString()
	return New(str)
}

// RightPad appends `i` whitespaces to the string
func (t Text) RightPad(i int) Text {
	return t.RightPadString(i, " ")
}

// RightPadString appends `i` times the given sequence to a string
func (t Text) RightPadString(i int, seq string) Text {
	return New(t.ToString() + strings.Repeat(seq, i))
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
