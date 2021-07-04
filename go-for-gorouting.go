package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// 10 3 10 10 10 10 10 10 10 10
		go func() {
			fmt.Println(i)
		}()

		// 9 3 0 5 1 4 7 6 2 8
		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}

	time.Sleep(1 * time.Millisecond)
}
