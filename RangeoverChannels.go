package main

import "fmt"

func main() {
	msgs := make(chan string, 2)

	msgs <- "hello"
	msgs <- "world"
	// msgs <- "world2"

	close(msgs)

	for msg := range msgs {
		fmt.Println(msg)
	}
}
