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
- [ ] delete
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
- [ ] ljust
- [x] lstrip
- [x] match
- [ ] next
- [ ] oct
- [ ] ord
- [ ] partition
- [ ] prepend
- [x] replace
- [x] reverse
- [ ] rindex
- [ ] rjust
- [ ] rpartition
- [x] rstrip
- [ ] scan
- [ ] scrub
- [ ] setbyte
- [ ] slice
- [x] split
- [ ] squeeze
- [x] start_with?
- [x] strip
- [x] sub
- [ ] succ
- [ ] sum
- [ ] swapcase
- [ ] to_c
- [ ] to_f
- [ ] to_i
- [ ] to_r
- [x] to_s
- [x] to_str
- [ ] to_sym
- [ ] tr
- [ ] tr_s
- [ ] unpack
- [x] upcase
- [ ] upto
- [ ] valid_encoding?