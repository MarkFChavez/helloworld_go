package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func definitive_swap(x, y string) (string, int) {
	return y, 24
}

func main() {
	a, x := definitive_swap("mark", "joel")
	fmt.Println(a, x)
}
