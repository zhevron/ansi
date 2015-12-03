ansi - ANSI escape sequence library
===================================

[![wercker status](https://app.wercker.com/status/387e7da9a8e93d6b23892db0ea3ecf51/s/master "wercker status")](https://app.wercker.com/project/bykey/387e7da9a8e93d6b23892db0ea3ecf51)
[![Coverage Status](https://coveralls.io/repos/zhevron/ansi/badge.svg?branch=master&service=github)](https://coveralls.io/github/zhevron/ansi?branch=master)
[![GoDoc](https://godoc.org/github.com/zhevron/ansi?status.svg)](https://godoc.org/github.com/zhevron/ansi)

**ansi** is an ANSI escape sequence library for [Google Go](https://golang.org).

## Colored strings

If the output supported ANSI escape codes, you can print colored strings using
the utility functions in this library.

```go
fmt.Println(ansi.Red("This string will be red", None))
```

You can also apply modifiers to the output strings to make the text bold, italic
or underlined.

```go
fmt.Println(ansi.Red("This string will be bold red", Bold))
fmt.Println(ansi.Red("This string will be bold and underlined red", Bold|Underline))
```

## License

**ansi** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
