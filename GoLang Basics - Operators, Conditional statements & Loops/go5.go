package main

import (
	"fmt"
)

func main() {
	const p = 10
	const q = "queen"
	fmt.Printf("%d \n", p)
	fmt.Printf("%s \n", q)
	p = 20
	q = "king"
	fmt.Printf("%d \n", p)
	fmt.Printf("%s \n", q)
}
