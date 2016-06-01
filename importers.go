package text

/*
 * Constructors / Import Utilities from various types
 */

import "strings"

// FromBytes does the casting for you
func FromBytes(b []byte) Text {
	return New(string(b))
}

// FromString is the same as New()
func FromString(s string) Text {
	return New(s)
}

// FromChars takes an []string slice of (one or more) characters
func FromChars(c []string) Text {
	return New(strings.Join(c, ""))
}

// New is the canonical constructor, same as FromString()
func New(s string) Text {
	return Text(s)
}
