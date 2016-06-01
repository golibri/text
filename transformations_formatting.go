package text

/*
 * transformations: chainable formatting helpers (immutable)
 */

import "strings"

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
