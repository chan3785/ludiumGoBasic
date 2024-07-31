package main

import (
	"fmt"
)

func greet() {
	fmt.Println("Hello, World!")
}

func add(a int, b int) int {
	return a + b
}

func main() {
	greet()
	result := add(4, 3)
	fmt.Println("Sum:", result)
}
