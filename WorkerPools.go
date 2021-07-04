package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "start job", j)

		time.Sleep(time.Second)

		fmt.Println("worker", id, "finish job", j)

		results <- j * 2
	}
}

func main() {
	const num = 5

	jobs := make(chan int, num)
	results := make(chan int, num)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= num; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 0; a < num; a++ {
		<-results
	}
}
