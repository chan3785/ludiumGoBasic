package main

import (
	"fmt"
)

func main() {
	for {
		var input string
		fmt.Print("Enter command: ")
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		} else if input == "hello" {
			fmt.Println("Hello, world!")
		} else if input == "even" {
			fmt.Println("Even numbers from 0 to 10:")
			for i := 0; i <= 5; i++ {
				fmt.Println(2 * i)
			}
		} else if input == "odd" {
			fmt.Println("Odd numbers from 1 to 10:")
			for i := 0; i < 5; i++ {
				fmt.Println(2*i + 1)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
