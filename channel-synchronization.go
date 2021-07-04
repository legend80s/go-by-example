package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")

	time.Sleep(time.Second)

	done <- true
}

func main() {
	fmt.Println("begin", time.Now())

	done := make(chan bool)

	go worker(done)

	<-done

	fmt.Println("done", time.Now())
}
