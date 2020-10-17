package main

import "fmt"

func main() {
	firstName := "Amit"
	userID, docId := 1, 10
	percentage := 75.0
	engineer := true

	fmt.Printf("%v \n ", firstName)
	fmt.Printf("%v \n ", percentage)
	fmt.Printf("%v \n ", userID)
	fmt.Printf("%v \n ", docId)
	fmt.Printf("%v \n ", engineer)

	fmt.Println()
	fmt.Printf("%T \n", firstName)
	fmt.Printf("%T \n", userID)
	fmt.Printf("%T \n", percentage)
	fmt.Printf("%T \n", engineer)
}
