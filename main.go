package main

import "fmt"

func main() {
	sum := add(4, 5)
	fmt.Println("Sum is", sum)
}

func add(x, y int) int {
	return x + y
}
