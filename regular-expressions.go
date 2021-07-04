package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var log = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	log(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	log(r.MatchString("peach"))

	log(r.FindString("peach punch"))
	log("FindAllString", r.FindAllString("peach punch pinch", -1))
	log("FindAllStringIndex", r.FindAllStringIndex("peach punch pinch", -1))

	loc := r.FindStringIndex("peach punch")
	log("peach punch"[loc[0]:loc[1]])

	log("FindStringSubmatch:", r.FindStringSubmatch("peach punch"))
	log("FindAllStringSubmatch:", r.FindAllStringSubmatch("peach punch", -1))
	log("FindStringSubmatchIndex:", r.FindStringSubmatchIndex("peach punch"))
	log("FindAllStringSubmatchIndex:", r.FindAllStringSubmatchIndex("peach punch", -1))

	log("Match", r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")

	log("r", r)

	log(r.ReplaceAllString("a peach", "<fruit>"))
	log(r.ReplaceAllStringFunc("a peach", func(s string) string {
		return s + " (fruit)"
	}))

	log(r.ReplaceAllStringFunc("a peach", strings.ToUpper))

	log(r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper))
	log(string(r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)))
}
