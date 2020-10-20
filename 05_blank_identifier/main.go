package main

import "fmt"

func main() {

	fruits := [3]string{"Apple", "Orange", "Mango"}

	// Range returns both the current index and value
	// but sometimes you may only want to use the value
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
