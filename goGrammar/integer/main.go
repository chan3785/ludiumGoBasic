package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = -459
	var b int32 = -2147483648
	var c int32 = 2147483647
	var d uint32 = 4294967295
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var e int8 = 127
	fmt.Println(e)
	e++
	fmt.Println(e)

	var f uint8 = 255
	fmt.Println(f)
	f++
	fmt.Println(f)

	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxUint32)
	fmt.Println(math.MaxInt64)

	var maxUint64 uint64 = math.MaxUint64
	fmt.Println(maxUint64)
}
