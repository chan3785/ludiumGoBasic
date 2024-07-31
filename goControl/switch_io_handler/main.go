package main

import (
	"fmt"
)

func main() {
	for {
		var input string
		fmt.Print("Enter command: ")
		fmt.Scanln(&input)
		switch {
		case input == "exit":
			fmt.Println("Exiting...")
		case input == "hello":
			fmt.Println("Hello, world!")
		case input == "even":
			fmt.Println("Even numbers from 0 to 10:")
			for i := 0; i <= 5; i++ {
				fmt.Println(2 * i)
			}
		case input == "odd":
			fmt.Println("Odd numbers from 1 to 10:")
			for i := 0; i < 5; i++ {
				fmt.Println(2*i + 1)
			}
		case input == "help":
			fmt.Println("Available commands:")
			fmt.Println("exit - Exit the program")
			fmt.Println("hello - Print 'Hello, world!'")
			fmt.Println("even - Print even numbers from 0 to 10")
			fmt.Println("odd - Print odd numbers from 1 to 10")
			fmt.Println("help - Show this help message")
		default:
			fmt.Println("Unknown command")
		}

		if input == "exit" {
			break
		}
	}
}
