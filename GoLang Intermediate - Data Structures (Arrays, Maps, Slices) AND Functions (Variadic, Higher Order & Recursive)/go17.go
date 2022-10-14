package main

import (
	"fmt"
)

func add2Num(a int, b int) int {
	c := a + b
	return c
}

func main() {
	fmt.Println(add2Num(5, 7))
	sum := add2Num(3, 4)
	fmt.Println(sum)
}
