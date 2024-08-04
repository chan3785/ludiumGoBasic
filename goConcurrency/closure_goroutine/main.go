package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)

	for i := 0; i < 5; i++ {
		go func(num int) {
			fmt.Printf("%d", num)
			done <- true
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-done
	}
}
