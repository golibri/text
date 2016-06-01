# Golibri/text

This package is a thin wrapper over the built-in `string` type and provides several methods for one would expect from strings. Transformations are chainable, convenience checks work
consistently and unicode characters are respected. The methods use Go's stdlib functions wherever possible, wrapping them in a programmer-friendly way.

Inspiration: Ruby's String class