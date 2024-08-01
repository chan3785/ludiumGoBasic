package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Person struct{}

func (s Person) Speak() string {
	return "Hello, my name is Alice"
}

type Dog struct{}

func (s Dog) Speak() string {
	return "Woof! My name is Buddy"
}

func Greet(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	var p Person
	var d Dog

	Greet(p)
	Greet(d)
}
