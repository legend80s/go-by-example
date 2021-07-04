package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

// non-declaration statement outside function body
// p := fmt.Println

func main() {
	p("hello")

	p("Contains", s.Contains("test", "es"))
	p("Count", s.Count("test", "t"))
	p("HasPrefix", s.HasPrefix("test", "t"))
	p("HasSuffix", s.HasSuffix("test", "t"))
	p("Index", s.Index("test", "e"))
	p("Join", s.Join([]string{"a", "b"}, "-"))
	p("Split", s.Split("a-b-c-d", "-"))
	p("Repeat", s.Repeat("a", 5))

	p("Replace", s.Replace("foo", "o", "0", -1))
	p("Replace", s.Replace("foo", "o", "0", 1))
	p("ToLower", s.ToLower("TEST"))
	p("ToUpper", s.ToUpper("test"))

	p()
	p("Len", len("hello"))
	p("Char", "hello"[1])
}
