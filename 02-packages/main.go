package main

import (
	"fmt"

	"github.com/amit-dhakad/learn-go-programming/02-packages/numbers"
	"github.com/amit-dhakad/learn-go-programming/02-packages/strings"
	"github.com/amit-dhakad/learn-go-programming/02-packages/strings/greeting"
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("hello world"))

}
