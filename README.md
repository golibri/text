# Golibri/text

This package is a thin wrapper over the built-in `string` type and provides several methods for one would expect from strings. Transformations are chainable, convenience checks work
consistently and unicode characters are respected. The methods use Go's stdlib functions wherever possible, wrapping them in a programmer-friendly way.

Inspiration: Ruby's String class:

- [x] ascii_only?
- [x] bytes
- [x] bytesize
- [x] capitalize
- [x] chars
- [x] count
- [x] delete
- [x] downcase
- [x] each_byte
- [x] each_char
- [x] each_line
- [x] empty?
- [x] end_with?
- [x] eql?
- [x] gsub
- [x] include?
- [ ] index
- [ ] insert
- [ ] inspect
- [ ] intern
- [x] length
- [x] lines
- [x] ljust
- [x] lstrip
- [x] match
- [x] replace
- [x] reverse
- [x] rjust
- [x] rstrip
- [ ] slice
- [x] split
- [x] start_with?
- [x] strip
- [x] sub
- [x] to_f
- [x] to_i
- [x] to_s
- [x] to_str
- [ ] tr
- [ ] tr_s
- [x] upcase