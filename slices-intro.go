package main

import "fmt"

var log = fmt.Println

func main() {
	b := [...]string{"Penn", "Teller"}

	log(b[0])
	log(b[1])
	log(len(b))
	log(cap(b))
	// log(b[-1]) // Compile error: invalid array index -1 (index must be non-negative)
	// log(b[2]) // Compile error: invalid array index 2 (out of bounds for 2-element array)

	sb := b[:]
	log(sb)

	log(sb[0])
	log(sb[1])
	log(len(sb))
	log(cap(sb))

	// s := make([]byte, 5, 4) // ./slices-intro.go:25:11: len larger than cap in make([]byte)
	s := make([]byte, 5, 6)
	log(s)
	log(len(s))
	log(cap(s))
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	total := m + len(data)

	if total > cap(slice) { // if necessary, rellocate
		// allocate double what's needed, for future growth
		newSlice := make([]byte, (total+1)*2)

		copy(newSlice, slice)

		slice = newSlice
	}

	slice = slice[0:total]

	copy(slice[m:total], data)

	return slice
}
