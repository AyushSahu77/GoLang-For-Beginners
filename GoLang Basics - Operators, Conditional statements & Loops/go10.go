package main

import (
	"fmt"
)

func main() {
	var a int = 12
	var b int = 25
	var p int = 212
	var q int = 212
	c := a & b
	d := a | b
	e := a ^ b
	f := p << 1
	g := q >> 1
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
