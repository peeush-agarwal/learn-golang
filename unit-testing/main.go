package main

import "fmt"

func main() {
	sum := add(1, 2)
	fmt.Printf("Sum of 1 and 2 is %v\n", sum)
}

func add(x int, y int) int {
	return x + y
}
