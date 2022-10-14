package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a string = "aaaa"
	var b int = 4
	var c bool = false
	var d float64 = 3.47
	var e string = "eeee"
	fmt.Printf("value of a %s and data type of a is %T \n", a, a)
	fmt.Printf("value of b %d and data type of b is %T \n", b, b)
	fmt.Printf("value of c %v and data type of c is %T \n", c, c)
	fmt.Printf("value of d %.2f and data type of d is %T \n", d, d)
	fmt.Printf("type of e is %v", reflect.TypeOf(e))
}
