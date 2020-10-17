package main

import (
	"fmt"

	"github.com/amit-dhakad/learn-go-programming/04_scope/math"
)

var a int = 10

func main() {
	b := 12
	math.PrintPi()
	fmt.Println(b)
	test()
}

func test() {
	// test don't have access of b
	// fmt.Println(b)

	fmt.Println(a)

}
