package main

import "fmt"

func main() {
	fmt.Println("1", isPrime(1))   // false
	fmt.Println("-1", isPrime(-1)) // false
	fmt.Println("2", isPrime(2))   // true
	fmt.Println("3", isPrime(3))   // true
	fmt.Println("4", isPrime(4))   // false
	fmt.Println("5", isPrime(5))   // true
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
