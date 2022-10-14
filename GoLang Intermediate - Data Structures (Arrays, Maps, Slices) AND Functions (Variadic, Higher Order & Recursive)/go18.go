package main

import (
	"fmt"
)

func operation1(a int, b int) (int, int) {
	c := a + b
	d := a - b
	return c, d
}

func operation2(e int, f int) (mul int, div int) {
	mul = e * f
	div = e / f
	return // only return as what to return is only written at declaration
}

func main() {
	sum, diff := operation1(7, 5)
	mul, div := operation2(12, 4)
	fmt.Println(sum)
	fmt.Println(diff)
	fmt.Println(mul)
	fmt.Println(div)
}
