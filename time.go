package main

import (
	"fmt"
	"time"
)

var log = fmt.Println

func main() {
	now := time.Now()
	log(now)
	log(now.Location())

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC,
	)

	log(then)
	log(then.Year())
	log(then.Month())
	log(then.Day())
	log(then.Hour())
	log(then.Minute())
	log(then.Second())
	log(then.Nanosecond())
	log(then.Location())

	log(now.Sub(then))
}
