package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	// messages := make(chan string, 1)
	signals := make(chan string)
	// messages <- "wow"

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case msg := <-signals:
		fmt.Println("received signal", msg)

	default:
		fmt.Println("no activity")
	}
}
