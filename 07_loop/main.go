package main

import "fmt"

func main() {

	i := 1

	for i < 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 100; j++ {

		if j%2 != 0 {
			continue
		}
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}
