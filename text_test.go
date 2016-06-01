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
	s := "gopher"
	equal(FromString(s).ToString(), s, "FromString() error", t)
	b := []byte("gopher")
	equal(FromBytes(b).ToBytes()[0], b[0], "FromBytes() error", t)
	c := []string{"g", "o", "p", "h", "e", "r"}
	equal(FromChars(c).ToChars()[0], c[0], "FromChars() error", t)
}

func TestMetrics(t *testing.T) {
	txt := Text("gopher")
	equal(txt.ByteSize(), 6, "ByteSize() error", t)
	equal(txt.Count("o"), 1, "ByteSize() error", t)
	equal(txt.Length(), 6, "ByteSize() error", t)
}

func TestTransformations(t *testing.T) {
	txt := Text("gopher")
	equal(txt.Capitalize(), Text("Gopher"), "Capitalize() error", t)
	equal(txt.ReplaceString("o", "a"), Text("gapher"), "ReplaceString() error", t)
	equal(txt.ReplacePattern("[oe]", "a"), Text("gaphar"), "ReplacePattern() error", t)
	equal(txt.Reverse(), Text("rehpog"), "Reverse() error", t)
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
