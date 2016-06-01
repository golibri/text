package text

import "testing"

var testWord = `Gopher`

func TestMethods(t *testing.T) {
	txt := Text("gopher")
	equal(len(txt.ToBytes()), 6, "Bytes() error", t)

	equal(txt.ToString(), "gopher", "ToString() error", t)
	equal(txt.ToChars()[2], "p", "Chars() error", t)
}

func TestChecks(t *testing.T) {
	txt := Text("gopher")
	truth(txt.EndsWith("er"), "EndsWith() truth", t)
	untruth(txt.EndsWith("stuff"), "EndsWith() untruth", t)
	truth(txt.IsASCIIOnly(), "IsAsciiOnly() truth", t)
	untruth(New("hören").IsASCIIOnly(), "IsAsciiOnly() untruth", t)
	truth(New("").IsEmpty(), "IsEmpty() truth", t)
	untruth(txt.IsEmpty(), "IsEmpty() untruth", t)
	truth(txt.StartsWith("go"), "StartsWith() truth", t)
	untruth(txt.StartsWith("stuff"), "StartsWith() untruth", t)
}

func TestImportersExporters(t *testing.T) {
	b := []byte("gopher")
	equal(FromBytes(b).ToBytes()[0], b[0], "FromBytes() error", t)
	c := []string{"g", "o", "p", "h", "e", "r"}
	equal(FromChars(c).ToChars()[0], c[0], "FromChars() error", t)
	l := "go\nis\r\ncool"
	equal(New(l).ToLines()[2], "cool", "ToLines() error", t)
	s := "gopher"
	equal(FromString(s).ToString(), s, "FromString() error", t)
}

func TestIterators(t *testing.T) {
	txt := New("Go\nis simple\nhöhö")

	bs := []byte{}
	txt.EachByte(func(b byte) {
		bs = append(bs, b)
	})
	equal(txt.ToBytes()[0], bs[0], "EachByte() error", t)

	cs := []string{}
	txt.EachChar(func(c string) {
		cs = append(cs, c)
	})
	equal(txt.ToChars()[0], cs[0], "EachChar() error", t)

	ls := []string{}
	txt.EachLine(func(l string) {
		ls = append(ls, l)
	})
	equal(txt.ToLines()[0], ls[0], "EachLine() error", t)
}

func TestMetrics(t *testing.T) {
	txt := Text("gopher")
	equal(txt.ByteSize(), 6, "ByteSize() error", t)
	equal(txt.Count("o"), 1, "Count() error", t)
	equal(txt.Length(), 6, "Length() error", t)
}

func TestTransformations(t *testing.T) {
	txt := Text("gopher")
	equal(txt.Capitalize(), Text("Gopher"), "Capitalize() error", t)
	equal(txt.ReplaceString("o", "a"), Text("gapher"), "ReplaceString() error", t)
	equal(txt.ReplacePattern("[oe]", "a"), Text("gaphar"), "ReplacePattern() error", t)
	equal(txt.Reverse(), Text("rehpog"), "Reverse() error", t)
	equal(txt.Capitalize().Downcase(), Text("gopher"), "Downcase() error", t)
	equal(txt.Upcase(), Text("GOPHER"), "Upcase() error", t)
}

/*
 * test helper functions to reduce boilerplate
 */

func truth(result bool, msg string, t *testing.T) {
	if result != true {
		t.Error(msg)
	}
}

func untruth(result bool, msg string, t *testing.T) {
	if result != false {
		t.Error(msg)
	}
}

func equal(a interface{}, b interface{}, msg string, t *testing.T) {
	if a != b {
		t.Errorf("%v :: equality Mismatch: %v != %v", msg, a, b)
	}
}
