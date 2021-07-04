package main

import (
	"fmt"
	"sort"
)

type byLen []string

func (s byLen) Len() int {
	return len(s)
}

func (s byLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLen(fruits))
	fmt.Println(fruits)

	ints := []int{5, 3, 4, 7, 8, 9}

	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})

	fmt.Println(ints)
}
