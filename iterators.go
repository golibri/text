package text

import "strings"

/*
 * iterators: apply functions incrementally
 */

// EachByte applies the given function on every byte of the string
func (t Text) EachByte(f func(byte)) {
	for _, b := range t.ToBytes() {
		f(b)
	}
}

// EachChar applies the given function on every char of the string
func (t Text) EachChar(f func(string)) {
	for _, c := range t.ToChars() {
		f(c)
	}
}

// EachLine applies the given function on every Line of the string
func (t Text) EachLine(f func(string)) {
	for _, l := range t.ToLines() {
		f(l)
	}
}

// MapBytes applies a given function on every byte directly
func (t Text) MapBytes(f func(byte) byte) Text {
	result := []byte{}
	for _, b := range t.ToBytes() {
		result = append(result, f(b))
	}
	return FromBytes(result)
}

// MapChars applies a given function on every char directly
func (t Text) MapChars(f func(string) string) Text {
	result := []string{}
	for _, c := range t.ToChars() {
		result = append(result, f(c))
	}
	return FromChars(result)
}

// MapLines applies a given function on every line directly
func (t Text) MapLines(f func(string) string) Text {
	result := []string{}
	for _, l := range t.ToLines() {
		result = append(result, f(l))
	}
	return New(strings.Join(result, "\n"))
}
