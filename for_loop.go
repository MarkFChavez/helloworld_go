package main

import "fmt"

func main() {
	//standard for loop
	for i := 1; i <= 10; i++ {
		fmt.Println("The current index is", i)
	}

	//shorten by
	x := 1
	for x < 100 {
		fmt.Println("Number is", x)
		x++
	}
}
