package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if t == v {
			return i
		}
	}

	return -1
}

func Includes(vs []string, t string) bool {
	return Index(vs, t) != -1
}

func Any(vs []string, predicate func(i int, t string) bool) bool {
	for i, v := range vs {
		if predicate(i, v) {
			return true
		}
	}

	return false
}

func All(vs []string, predicate func(i int, t string) bool) bool {
	for i, v := range vs {
		if !predicate(i, v) {
			return false
		}
	}

	return true
}

func Filter(vs []string, predicate func(index int, value string) bool) []string {
	// results := []string{}
	var results []string
	// fmt.Println("results slice expected to be nil:", results)
	fmt.Println("results slice expected to be nil:", results == nil)

	for i, v := range vs {
		if predicate(i, v) {
			results = append(results, v)
		}
	}

	return results
}

func Map(vs []string, mapper func(value string, index int) string) []string {
	var results = []string{}

	for i, v := range vs {
		results = append(results, mapper(v, i))
	}

	return results
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "apple"))    // 1
	fmt.Println(Includes(strs, "grape")) // false
	fmt.Println(Includes(strs, "plum"))  // true
	fmt.Println(Any(strs, func(i int, t string) bool {
		return strings.HasPrefix(t, "p")
	})) // true
	fmt.Println(Any(strs, func(i int, t string) bool {
		return strings.HasSuffix(t, "p")
	})) // false
	fmt.Println("Filtered with e", strs, Filter(strs, func(index int, value string) bool {
		return strings.Contains(value, "e")
	})) // ["peach", "apple", "pear"]
	fmt.Println(Map(strs, func(value string, index int) string {
		return strings.ToUpper(value)
	}))
	// ["PEACH", "APPLE", "PEAR", "PLUM"]
}
