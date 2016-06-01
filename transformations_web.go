package text

/*
 * transformations: chainable string transformations for web interactions
 */

import "html"

// EscapeHTML translates `<` to `&lt;`, and so on
func (t Text) EscapeHTML() Text {
	return New(html.EscapeString(t.ToString()))
}

// UnescapeHTML translates `&lt;` to `<`, and so on
func (t Text) UnescapeHTML() Text {
	return New(html.UnescapeString(t.ToString()))
}
