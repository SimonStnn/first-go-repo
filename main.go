package main

import "fmt"

// First function
func add(x int, y int) int {
	return x + y
}

func main() {
	// Print "Hello, World!" to the terminal.
	fmt.Println("Hello, World!")

	// Print the result of the add function.
	fmt.Println(add(42, 13))
}