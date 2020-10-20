package main

import "fmt"

const a string = "hello"

func main() {
	const b = 10
	const AB = 5 // don't use A_B

	fmt.Println("A- ", AB)
	fmt.Println("a- ", a)
	fmt.Println("b - ", b)
}
