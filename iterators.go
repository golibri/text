package text

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
