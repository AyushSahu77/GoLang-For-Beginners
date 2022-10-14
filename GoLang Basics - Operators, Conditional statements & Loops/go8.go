package main

import (
	"fmt"
)

func main() {
	var a int = 50
	fmt.Println((a < 100) && (a < 200))
	fmt.Println((a < 100) && (a > 200))
	fmt.Println("\n")
	fmt.Println((a < 100) || (a < 200))
	fmt.Println((a < 100) || (a > 200))
	fmt.Println((a > 100) || (a > 200))
	fmt.Println("\n")
	fmt.Println(!(a > 100))
	fmt.Println(!(a < 100))
	fmt.Println(!(true))
	fmt.Println(!(false))
}
