package main

import (
	"fmt"
)

func main() {
	var a int = 7
	var b int = a
	var c int = 10
	c += a
	var d int = 20
	d -= a
	var e int = 30
	e *= a
	var f int = 40
	f /= a
	var g int = 50
	g %= a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
