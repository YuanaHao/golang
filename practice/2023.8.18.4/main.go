package main

import (
	"fmt"
)

func sum1(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func main() {
	ret1 := sum1(0)
	ret2 := sum1(10, 20)
	ret3 := sum1(10, 20, 30, 40)
	fmt.Println(ret1, ret2, ret3)
}
