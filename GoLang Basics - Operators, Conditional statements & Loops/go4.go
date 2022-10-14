package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 4
	var b float64 = float64(a)
	var c int = 7
	var s1 string = strconv.Itoa(c)
	var s2 string = "Hello"
	//var d int = strconv.Atoi(s2)
	fmt.Printf("%.2f \n", b)
	fmt.Printf("%q \n", s1)
	i, err := strconv.Atoi(s2)
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T \n", err, err)
	//fmt.Printf("%d \n", d)
}
