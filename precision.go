package main

import "fmt"

func main() {
	a1 := 0.1
	b1 := 0.2
	c1 := 0.3
	fmt.Println("float64")
	fmt.Println("0.1 + 0.2 === 0.3 ?", a1+b1 == 0.3)
	fmt.Println("0.1 + 0.2 === c ?", a1+b1 == c1)
	fmt.Println("0.1 + 0.2 = ?", a1+b1)

	fmt.Println("float32")
	var a float32 = 0.1
	var b float32 = 0.2
	var c float32 = 0.3
	fmt.Println("0.1 + 0.2 === 0.3 ?", a+b == 0.3)
	fmt.Println("0.1 + 0.2 === c ?", a+b == c)
	fmt.Println("0.1 + 0.2 = ?", a+b)

	fmt.Println("0.1 + 0.2 === 0.3 ?", 0.1+0.2 == 0.3)
	fmt.Println("0.1 + 0.2 = ?", 0.1+0.2)
}
