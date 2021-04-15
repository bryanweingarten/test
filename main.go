package main

import "fmt"

func main() {
	sum := add(2, 3)
	fmt.Println("Sum is", sum)
}

func add(x, y int) int {
	return x + y
}
