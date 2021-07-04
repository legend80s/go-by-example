package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C

	fmt.Println("Timer 1 fired after", 2*time.Second)

	// cancel a timer1

	timer2 := time.NewTimer(1 * time.Second)

	go func() {
		<-timer2.C

		fmt.Println("Timer 2 fired after", 1*time.Second)
	}()

	// stop := timer2.Stop()

	// if stop {
	// 	fmt.Println("Timer 2 stopped")
	// }

	time.Sleep(2 * time.Second)
}
