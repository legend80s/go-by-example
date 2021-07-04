package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs

			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")

				done <- true

				return
			}
		}
	}()

	for i := 1; i <= 10; i++ {
		jobs <- i
		fmt.Println("send job", i)
	}

	close(jobs)

	fmt.Println("send all jobs")

	// <-done
}
