package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "hell world" }()

	msg := <-messages

	fmt.Println(msg)
}
