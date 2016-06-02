[![Build Status](https://travis-ci.org/golibri/text.svg?branch=master)](https://travis-ci.org/golibri/text)
[![GoDoc](https://godoc.org/github.com/golibri/text?status.svg)](https://godoc.org/github.com/golibri/text)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# Golibri/text

This package is a thin wrapper over the built-in `string` type and provides several methods for one would expect from strings. Transformations are chainable, convenience checks work
consistently and unicode characters are respected. The methods use Go's stdlib functions wherever possible, wrapping them in a programmer-friendly way. Inspired by Ruby's String class

## Installation

`go get -u github.com/golibri/text`

## Usage examples

1. Create a text object, typically from a string:

  ````go
  txt := text.New(" hello world! ") // or: text.FromString("...")
  // also possible: text.FromBytes(...)
  ````

2. Check some properties of the text:

  ````go
  txt.IsASCIIOnly() // true
  txt.IsEmpty() // false
  txt.StartsWith(" hello") // true
  ````

3. Get some metrics from the text:

  ````go
  txt.Length() // 13
  txt.Count("l") // 3
  ````

4. Transform the text with chainable methods:

  ````go
  result := txt.
    ReplaceString("world", "golang").
    Capitalize().
    Strip() // text.Text("Hello Golang")
  ````

5. Export the text back to a usual string:

  ````go
  result.ToString() // "Hello Golang"
  ````

## Caveats

This package aims to provide a convenient programming interface and is not optimized for the best possible performance, there are memory allocations, especially for the transformation methods.

For a good overview about all the available methods, checkout the godoc documentation!

## License

LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)