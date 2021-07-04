package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating", p)

	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}

	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing", f)

	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		os.Exit(1)
	}
}

func writeFile(f *os.File) {
	fmt.Println("writing", f)

	content := "data"

	fmt.Fprintln(f, content)

	fmt.Println("written with content:", content)
}
