package main

import "fmt"

func add(x, y int) (int) {
  return x + y
}

func main() {
  fmt.Println("The sum of 10 and 20 is", add(10, 20))
}
