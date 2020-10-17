package main

import "fmt"

// User user type
type User struct {
	firstName string
	lastName  string
	id        int
}

  var outsideFunc string ="Test"
func main() {

	var str string
	var flt float64
	var bl bool
	var user *User


	var test, per int = 100, 50

	fmt.Printf("%v \n", outsideFunc)
	fmt.Printf("%v \n", test)
	fmt.Printf("%v \n", per)
	fmt.Printf("%v \n", str)
	fmt.Printf("%v \n", flt)
	fmt.Printf("%v \n", bl)
	fmt.Printf("%v \n", user)

}
