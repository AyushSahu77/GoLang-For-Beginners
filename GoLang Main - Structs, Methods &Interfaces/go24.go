package main

import (
	"fmt"
)

func modifyPassByVal(a int) {
	a = +10
}

func modifyPassByRef(s *string) {
	*s = "Hi"
}

func main() {
	a := 4
	s := "Hello"
	fmt.Println(a)
	modifyPassByVal(a)
	fmt.Println(a) // value remains unchanged as this is pass by value
	fmt.Println(s)
	modifyPassByRef(&s)
	fmt.Println(s) // value is modified (from hello to hi) as this is pass by reference

	// MAPS AND SLICES ARE PASSED BY REFERENCE BY DEFAULT AS slice IS A REFERENCE TO AN UNDERLYING ARRAY AND maps ARE REFERENCED KEY-VALUE PAIRS
}
