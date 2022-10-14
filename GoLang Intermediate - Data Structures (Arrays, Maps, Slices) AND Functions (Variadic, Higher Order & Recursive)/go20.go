package main

import (
	"fmt"
)

var fact int = 1

func factorialSimple(a int) int {

	for i := a; i > 0; i-- {
		fact = fact * i
	}
	fmt.Printf("Factorial of %d is %d by simple func \n", a, fact)
	return fact
}

func factorialRecursive(b int) int {
	if b == 1 {
		return 1
	}
	return b * factorialRecursive(b-1)
}

func main() {
	factorialSimple(5)
	n := 5
	result := factorialRecursive(n)
	fmt.Printf("Factorial of %d is %d by recursive func \n", n, result)

	x := func(l int, b int) int { // anonymous function as it has no name and is directly used
		return l * b
	}
	fmt.Printf("%T \n", x) // type of x is func(int,int) int
	fmt.Println(x(20, 30)) // 600

	y := func(l int, b int) int {
		return l * b
	}(40, 50) // arguement are directly passed into the func

	fmt.Printf("%T \n", y) // type of x is int
	fmt.Println(y)         // 2000
}
