package main

import "fmt"

func main() {
	messages := make(chan string, 1)

	messages <- "hello"
	// fatal error: all goroutines are asleep - deadlock!
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-messages)
}
