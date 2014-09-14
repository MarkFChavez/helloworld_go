package main

import "fmt"

type Model struct {
	x, y int
}

func main() {
	t := Model{x: 20, y: 10}
	fmt.Println(t)
}
