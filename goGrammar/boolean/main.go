package main

import (
	"fmt"
)

func main() {
	var b1 bool = true
	var b2 bool = false

	fmt.Println(b1)
	fmt.Println(b2)

	var b3 bool
	fmt.Println(b3)

	a := 10
	b := 20

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)

	c := false
	d := true

	fmt.Println(c && d)
	fmt.Println(c || d)
	fmt.Println(!c)

}
